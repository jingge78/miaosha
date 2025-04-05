package model

import (
	"miaosha-jjl/common/global"
	"time"
)

type UserIntegralog struct {
	Id            string    `gorm:"column:id;type:varchar(50);primaryKey;not null;" json:"id"`
	UserId        int64     `gorm:"column:user_id;type:bigint;comment:用户id;not null;" json:"user_id"`                                                     // 用户id
	IntegralType  string    `gorm:"column:integral_type;type:varchar(50);comment:积分类型 1.签到 2.连续签到 3.福利任务 4.每日任务 5.补签;default:NULL;" json:"integral_type"` // 积分类型 1.签到 2.连续签到 3.福利任务 4.每日任务 5.补签
	Integral      string    `gorm:"column:integral;type:varchar(50);comment:积分;default:0;" json:"integral"`                                               // 积分
	Bak           string    `gorm:"column:bak;type:varchar(100);comment: 积分补充文案;default:NULL;" json:"bak"`                                                //  积分补充文案
	OperationTime time.Time `gorm:"column:operation_time;type:datetime(3);comment:操作时间(签到和补签的具体日期);default:NULL;" json:"operation_time"`                  // 操作时间(签到和补签的具体日期)
	CreateTime    time.Time `gorm:"column:create_time;type:datetime(3);comment:创建时间;default:NULL;" json:"create_time"`                                    // 创建时间
}

//注册用户

func (u *UserIntegralog) UserIntegralog() error {
	return global.DB.Debug().Table("UserId").Create(&u).Error
}

// 查询

func (u *UserIntegralog) UserIntegralogCx(Integral string) error {
	return global.DB.Table("Integral").Where("Integral=?", Integral).Limit(1).Find(&u).Error
}
