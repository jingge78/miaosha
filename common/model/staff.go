package model

import "miaosha-jjl/common/global"

type EbSystemStoreStaff struct {
	Id           uint32 `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	Uid          uint32 `gorm:"column:uid;type:int UNSIGNED;comment:微信用户id;not null;" json:"uid"`                    // 微信用户id
	Avatar       string `gorm:"column:avatar;type:varchar(255);comment:店员头像;not null;" json:"avatar"`                // 店员头像
	StoreId      uint32 `gorm:"column:store_id;type:int;comment:门店id;not null;" json:"store_id"`                     // 门店id
	StaffName    string `gorm:"column:staff_name;type:varchar(64);comment:店员名称;" json:"staff_name"`                  // 店员名称
	Phone        string `gorm:"column:phone;type:char(15);comment:手机号码;default:NULL;" json:"phone"`                  // 手机号码
	VerifyStatus uint32 `gorm:"column:verify_status;type:int;comment:核销开关;not null;default:0;" json:"verify_status"` // 核销开关
	Status       uint32 `gorm:"column:status;type:int;comment:状态;default:1;" json:"status"`                          // 状态
	AddTime      uint32 `gorm:"column:add_time;type:int;comment:添加时间;default:NULL;" json:"add_time"`                 // 添加时间
}

func (e *EbSystemStoreStaff) StaffIntegralog() error {
	return global.DB.Debug().Table("Uid").Create(&e).Error
}
