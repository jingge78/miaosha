package model

import "miaosha-jjl/common/global"

type UserEnter struct {
	Id           uint32 `gorm:"column:id;type:int UNSIGNED;comment:商户申请ID;primaryKey;not null;" json:"id"`                    // 商户申请ID
	Uid          uint32 `gorm:"column:uid;type:int UNSIGNED;comment:用户ID;not null;default:0;" json:"uid"`                     // 用户ID
	Province     string `gorm:"column:province;type:varchar(32);comment:商户所在省;not null;" json:"province"`                     // 商户所在省
	City         string `gorm:"column:city;type:varchar(32);comment:商户所在市;not null;" json:"city"`                             // 商户所在市
	District     string `gorm:"column:district;type:varchar(32);comment:商户所在区;not null;" json:"district"`                     // 商户所在区
	Address      string `gorm:"column:address;type:varchar(256);comment:商户详细地址;not null;" json:"address"`                     // 商户详细地址
	MerchantName string `gorm:"column:merchant_name;type:varchar(256);comment:商户名称;not null;" json:"merchant_name"`           // 商户名称
	LinkTel      string `gorm:"column:link_tel;type:varchar(16);comment:商户电话;not null;" json:"link_tel"`                      // 商户电话
	Charter      string `gorm:"column:charter;type:varchar(512);comment:商户证书;not null;" json:"charter"`                       // 商户证书
	AddTime      uint32 `gorm:"column:add_time;type:int UNSIGNED;comment:添加时间;not null;default:0;" json:"add_time"`           // 添加时间
	ApplyTime    uint32 `gorm:"column:apply_time;type:int UNSIGNED;comment:审核时间;not null;default:0;" json:"apply_time"`       // 审核时间
	SuccessTime  int32  `gorm:"column:success_time;type:int;comment:通过时间;not null;" json:"success_time"`                      // 通过时间
	FailMessage  string `gorm:"column:fail_message;type:varchar(256);comment:未通过原因;not null;" json:"fail_message"`            // 未通过原因
	FailTime     uint32 `gorm:"column:fail_time;type:int UNSIGNED;comment:未通过时间;not null;default:0;" json:"fail_time"`        // 未通过时间
	Status       int8   `gorm:"column:status;type:tinyint;comment:-1 审核未通过 0未审核 1审核通过;not null;default:0;" json:"status"`     // -1 审核未通过 0未审核 1审核通过
	IsLock       uint8  `gorm:"column:is_lock;type:tinyint UNSIGNED;comment:0 = 开启 1= 关闭;not null;default:0;" json:"is_lock"` // 0 = 开启 1= 关闭
	IsDel        uint8  `gorm:"column:is_del;type:tinyint UNSIGNED;comment:是否删除;not null;default:0;" json:"is_del"`           // 是否删除
}

func (e *UserEnter) AddUserEnter() error {
	return global.DB.Create(&e).Error
}
