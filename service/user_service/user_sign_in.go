package user_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"miaosha-jjl/common/global"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/user"
	"time"
)

/*
 用户签到功能
				签到逻辑
			1.签到检查
		2.连续签到计算
	3.积分计算：
				    补充
				    签到日历查询
					积分过期
*/

func (s *ServerUser) SignIn(ctx context.Context, req *user.SignInRequest) (resp *user.SignInResponse, err error) {

	// 如果传入了SignDate，支持自定义签到日期（方便测试）
	var signDate time.Time
	if req.SignDate != "" {
		signDate, err = time.Parse("2006-01-02", req.SignDate) //测试的话格式(2025-03-28)
		if err != nil {
			return nil, fmt.Errorf("无效的签到日期格式")
		}
	} else {
		signDate = time.Now() // 默认使用当前日期
	}

	today := signDate.Format("2006-01-02")
	//1.检查今天是否已经签到
	todaykey := fmt.Sprintf("sign:user:%d:%s", req.UserId, today)
	offset := signDate.Day() - 1 // 位图的偏移量从0开始
	bit, err := global.Rdb.GetBit(global.CTX, todaykey, int64(offset)).Result()
	if err != nil {
		return nil, err
	}
	if bit == 1 {
		return nil, fmt.Errorf("今天签到了")
	}

	//2.检查昨天是否已经签到，计算连续签到天数
	consecutiveDays := 1 // 默认连续1天
	yesterday := signDate.AddDate(0, 0, -1).Format("2006-01-02")
	yesterdayKey := fmt.Sprintf("sign:user:%d:%s", req.UserId, yesterday)
	//检查昨天的签到是否还存在
	exists, err := global.Rdb.Exists(global.CTX, yesterdayKey).Result()
	if err != nil {
		return nil, fmt.Errorf("昨天的签到不存在")
	}
	// 如果昨天有签到，获取连续签到天数
	if exists > 0 {
		consecutiveKey := fmt.Sprintf("sign:consecutive:%d", req.UserId)
		days, err := global.Rdb.Get(global.CTX, consecutiveKey).Int()
		if err != nil && err != redis.Nil {
			return nil, err
		}
		if days > 0 {
			consecutiveDays = days + 1
		}
	}
	// 3. 计算本次签到应得积分
	points := consecutiveDays // 第N天连续签到得N分
	// 4. 开启事务处理
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 5. 更新用户总积分
	var userIntegral model.UserIntegral
	if err := tx.Where("user_id = ?", req.UserId).First(&userIntegral).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果没有记录，创建新记录
			userIntegral = model.UserIntegral{
				ID:            uuid.New().String(),
				UserID:        int(req.UserId),
				Integral:      points,
				IntegralTotal: points,
				CreateTime:    signDate,
				UpdateTime:    signDate,
			}
			if err := tx.Create(&userIntegral).Error; err != nil {
				tx.Rollback()
				return nil, fmt.Errorf("创建积分记录失败")
			}
		} else {
			tx.Rollback()
			return nil, err
		}
	} else {
		// 更新现有记录
		if err := tx.Model(&userIntegral).Updates(map[string]interface{}{
			"integral":       gorm.Expr("integral + ?", points),
			"integral_total": gorm.Expr("integral_total + ?", points),
			"update_time":    signDate,
		}).Error; err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("更新积分失败")
		}
	}

	// 6. 创建积分流水记录
	integralLog := model.UserIntegralLog{
		ID:            uuid.New().String(),
		UserID:        int(req.UserId),
		IntegralType:  model.IntegralTypeContinuous, // 连续签到类型
		Integral:      points,
		Bak:           fmt.Sprintf("连续签到%d天", consecutiveDays),
		OperationTime: signDate,
		CreateTime:    signDate,
	}
	if err := tx.Create(&integralLog).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("创建积分流水失败")
	}

	// 7. 更新Redis中的签到状态和连续签到天数
	consecutiveKey := fmt.Sprintf("sign:consecutive:%d", req.UserId)
	pipe := global.Rdb.TxPipeline()
	pipe.SetBit(global.CTX, todaykey, int64(offset), 1)
	pipe.Set(global.CTX, consecutiveKey, consecutiveDays, 30*24*time.Hour) // 保留30天
	if _, err := pipe.Exec(global.CTX); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("更新签到状态失败")
	}

	// 8. 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("提交事务失败")
	}

	return &user.SignInResponse{
		Message: fmt.Sprintf("签到成功，连续签到%d天", consecutiveDays),
		Points:  int32(points),
	}, nil
}
