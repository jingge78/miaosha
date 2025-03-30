package user_service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"miaosha-jjl/common/global"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/user"
	"time"
)

func (s *ServerUser) MakeupSignIn(ctx context.Context, req *user.MakeupSignInRequest) (*user.MakeupSignInResponse, error) {
	// 1. 解析补签日期
	makeupDate, err := time.Parse("2006-01-02", req.SignDate)
	if err != nil {
		return nil, fmt.Errorf("无效的补签日期格式")
	}

	// 2. 检验补签日期 是不是在一周之内的
	if time.Since(makeupDate) > 7*24*time.Hour {
		return nil, fmt.Errorf("只能补签过去7天内的签到")
	}

	// 3. 检查是否已签到
	dateKey := fmt.Sprintf("sign:user:%d:%s", req.UserId, makeupDate.Format("2006-01-02"))
	offset := makeupDate.Day() - 1
	bit, err := global.Rdb.GetBit(ctx, dateKey, int64(offset)).Result()
	if err != nil {
		return nil, err
	}
	if bit == 1 {
		return nil, fmt.Errorf("该日期已签到，无需补签")
	}

	// 4. 检查用户是否有补签卡
	var makeupCard model.UserMakeupCard
	if err := global.DB.Where("user_id = ?", req.UserId).First(&makeupCard).Error; err != nil {
		return nil, fmt.Errorf("没有可用的补签卡")
	}
	if makeupCard.CardCount <= 0 {
		return nil, fmt.Errorf("没有可用的补签卡")
	}

	// 5. 计算积分（补签固定得1分）
	points := 1

	// 6. 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 7. 扣除补签卡
	if err := tx.Model(&model.UserMakeupCard{}).
		Where("user_id = ? AND cardCount > 0", req.UserId).
		Update("cardCount", gorm.Expr("cardCount - 1")).
		Update("update_time", time.Now()).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("扣除补签卡失败")
	}

	// 8. 更新用户积分
	if err := tx.Model(&model.UserIntegral{}).
		Where("user_id = ?", req.UserId).
		Updates(map[string]interface{}{
			"integral":       gorm.Expr("integral + ?", points),
			"integral_total": gorm.Expr("integral_total + ?", points),
			"update_time":    time.Now(),
		}).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("更新积分失败")
	}

	// 9. 创建积分流水记录
	integralLog := model.UserIntegralLog{
		ID:            uuid.New().String(),
		UserID:        int(req.UserId),
		IntegralType:  model.IntegralTypeReplenish, // 补签类型
		Integral:      points,
		Bak:           "补签获得",
		OperationTime: makeupDate,
		CreateTime:    time.Now(),
	}
	if err := tx.Create(&integralLog).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("创建积分流水失败")
	}

	// 10. 更新Redis签到状态（但不更新连续签到）
	_, err = global.Rdb.SetBit(ctx, dateKey, int64(offset), 1).Result()
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("更新签到状态失败")
	}

	// 11. 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("提交事务失败")
	}

	return &user.MakeupSignInResponse{
		Success: true,
		Message: "补签成功",
		Points:  int32(points),
	}, nil
}
