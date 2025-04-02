package model

import "miaosha-jjl/common/global"

type Coupon struct {
	Id          uint32  `gorm:"column:id;type:int UNSIGNED;comment:优惠券表ID;primaryKey;not null;" json:"id"`                                          // 优惠券表ID
	Title       string  `gorm:"column:title;type:varchar(64);comment:优惠券名称;not null;" json:"title"`                                                 // 优惠券名称
	Integral    uint32  `gorm:"column:integral;type:int UNSIGNED;comment:兑换消耗积分值;not null;default:0;" json:"integral"`                              // 兑换消耗积分值
	CouponPrice float64 `gorm:"column:coupon_price;type:decimal(8, 2) UNSIGNED;comment:兑换的优惠券面值;not null;default:0.00;" json:"coupon_price"`        // 兑换的优惠券面值
	UseMinPrice float64 `gorm:"column:use_min_price;type:decimal(8, 2) UNSIGNED;comment:最低消费多少金额可用优惠券;not null;default:0.00;" json:"use_min_price"` // 最低消费多少金额可用优惠券
	CouponTime  uint32  `gorm:"column:coupon_time;type:int UNSIGNED;comment:优惠券有效期限（单位：天）;not null;default:0;" json:"coupon_time"`                  // 优惠券有效期限（单位：天）
	Sort        uint32  `gorm:"column:sort;type:int UNSIGNED;comment:排序;not null;default:1;" json:"sort"`                                           // 排序
	Status      uint8   `gorm:"column:status;type:tinyint UNSIGNED;comment:状态（0：关闭，1：开启）;not null;default:0;" json:"status"`                        // 状态（0：关闭，1：开启）
	AddTime     uint32  `gorm:"column:add_time;type:int UNSIGNED;comment:兑换项目添加时间;not null;" json:"add_time"`                                       // 兑换项目添加时间
	IsDel       uint8   `gorm:"column:is_del;type:tinyint UNSIGNED;comment:是否删除;not null;default:0;" json:"is_del"`                                 // 是否删除
	ProductId   string  `gorm:"column:product_id;type:varchar(64);comment:所属商品id;not null;" json:"product_id"`                                      // 所属商品id
	CategoryId  int32   `gorm:"column:category_id;type:int;comment:分类id;not null;default:0;" json:"category_id"`                                    // 分类id
	Type        int8    `gorm:"column:type;type:tinyint;comment:优惠券类型 0-通用 1-品类券 2-商品券;not null;default:0;" json:"type"`                            // 优惠券类型 0-通用 1-品类券 2-商品券
}

func (c *Coupon) AddCoupon() error {
	return global.DB.Debug().Table("coupon").Create(&c).Error
}
