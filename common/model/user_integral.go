package model

import "time"

// UserIntegral 用户积分总表
type UserIntegral struct {
	ID            string    `gorm:"column:id;primaryKey;type:varchar(50)" json:"id"`      // id
	UserID        int       `gorm:"column:user_id;not null" json:"userId"`                // 用户id
	Integral      int       `gorm:"column:integral;default:0" json:"integral"`            // 当前积分
	IntegralTotal int       `gorm:"column:integral_total;default:0" json:"integralTotal"` // 累计积分
	CreateTime    time.Time `gorm:"column:create_time;autoUpdateTime" json:"createTime"`  // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`  // 修改时间
}

func (UserIntegral) TableName() string {
	return "user_integral"
}

// UserIntegralLog 用户积分流水表
type UserIntegralLog struct {
	ID            string    `gorm:"column:id;primaryKey;type:varchar(50)" json:"id"`      // id
	UserID        int       `gorm:"column:user_id;not null" json:"userId"`                // 用户id
	IntegralType  int       `gorm:"column:integral_type" json:"integralType"`             // 积分类型 1.签到 2.连续签到 3.福利任务 4.每日任务 5.补签
	Integral      int       `gorm:"column:integral;default:0" json:"integral"`            // 积分
	Bak           string    `gorm:"column:bak;type:varchar(100)" json:"bak"`              // 积分补充文案
	OperationTime time.Time `gorm:"column:operation_time;type:date" json:"operationTime"` // 操作时间(签到和补签的具体日期)
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`                 // 创建时间
}

func (UserIntegralLog) TableName() string {
	return "user_integral_log"
}

// 积分类型常量
const (
	IntegralTypeSignIn     = 1 // 签到
	IntegralTypeContinuous = 2 // 连续签到
	IntegralTypeWelfare    = 3 // 福利任务
	IntegralTypeDaily      = 4 // 每日任务
	IntegralTypeReplenish  = 5 // 补签
)
