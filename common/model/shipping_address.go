package model

import (
	"miaosha-jjl/common/global"
	"time"
)

type ShippingAddress struct {
	Id              uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:收货地址ID;primaryKey;not null;" json:"id"` // 收货地址ID
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime(3);not null;" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
	DeletedAt       time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	UserId          uint64    `gorm:"column:user_id;type:bigint UNSIGNED;comment:用户ID;not null;" json:"user_id"`                      // 用户ID
	RecipientName   string    `gorm:"column:recipient_name;type:varchar(20);comment:收货人姓名;not null;" json:"recipient_name"`           // 收货人姓名
	PhoneNumber     string    `gorm:"column:phone_number;type:char(11);comment:联系电话;not null;" json:"phone_number"`                   // 联系电话
	Province        string    `gorm:"column:province;type:varchar(32);comment:省份;not null;" json:"province"`                          // 省份
	City            string    `gorm:"column:city;type:varchar(32);comment:城市;not null;" json:"city"`                                  // 城市
	District        string    `gorm:"column:district;type:varchar(32);comment:区/县;not null;" json:"district"`                         // 区/县
	DetailedAddress string    `gorm:"column:detailed_address;type:varchar(32);comment:详细地址;not null;" json:"detailed_address"`        // 详细地址
	IsDefault       uint64    `gorm:"column:is_default;type:bigint UNSIGNED;comment:是否为默认地址，0 表示否，1 表示是;not null;" json:"is_default"` // 是否为默认地址，0 表示否，1 表示是
}

func (a *ShippingAddress) AddShippingAddress() error {
	return global.DB.Debug().Table("shipping_address").Create(&a).Error
}
