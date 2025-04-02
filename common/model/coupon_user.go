package model

type CouponUser struct {
	Id          int32   `gorm:"column:id;type:int;comment:优惠券发放记录id;primaryKey;not null;" json:"id"`                                                // 优惠券发放记录id
	Uid         uint32  `gorm:"column:uid;type:int UNSIGNED;comment:优惠券所属用户;not null;default:0;" json:"uid"`                                        // 优惠券所属用户
	CouponTitle string  `gorm:"column:coupon_title;type:varchar(32);comment:优惠券名称;not null;" json:"coupon_title"`                                   // 优惠券名称
	CouponPrice float64 `gorm:"column:coupon_price;type:decimal(8, 2) UNSIGNED;comment:优惠券的面值;not null;default:0.00;" json:"coupon_price"`          // 优惠券的面值
	UseMinPrice float64 `gorm:"column:use_min_price;type:decimal(8, 2) UNSIGNED;comment:最低消费多少金额可用优惠券;not null;default:0.00;" json:"use_min_price"` // 最低消费多少金额可用优惠券
	AddTime     uint32  `gorm:"column:add_time;type:int UNSIGNED;comment:优惠券创建时间;not null;" json:"add_time"`                                        // 优惠券创建时间
	EndTime     uint32  `gorm:"column:end_time;type:int UNSIGNED;comment:优惠券结束时间;not null;" json:"end_time"`                                        // 优惠券结束时间
	Type        string  `gorm:"column:type;type:varchar(32);comment:获取方式;not null;default:send;" json:"type"`                                       // 获取方式
	Status      int8    `gorm:"column:status;type:tinyint(1);comment:状态（0：未使用，1：已使用, 2:已过期）;not null;default:0;" json:"status"`                     // 状态（0：未使用，1：已使用, 2:已过期）
}
