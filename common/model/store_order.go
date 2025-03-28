package model

import (
	"gorm.io/gorm"
	"miaosha-jjl/common/global"
)

type StoreOrder struct {
	gorm.Model
	OrderId      string  `gorm:"column:order_id;type:varchar(100);comment:订单号;not null;" json:"order_id"`                                                          // 订单号
	Title        string  `gorm:"column:title;type:varchar(255);comment:标题;not null;" json:"title"`                                                                 // 标题
	Uid          uint32  `gorm:"column:uid;type:int UNSIGNED;comment:用户id;not null;" json:"uid"`                                                                   // 用户id
	RealName     string  `gorm:"column:real_name;type:varchar(32);comment:用户姓名;not null;" json:"real_name"`                                                        // 用户姓名
	UserPhone    string  `gorm:"column:user_phone;type:varchar(18);comment:用户电话;not null;" json:"user_phone"`                                                      // 用户电话
	UserAddress  string  `gorm:"column:user_address;type:varchar(100);comment:详细地址;not null;" json:"user_address"`                                                 // 详细地址
	CartId       string  `gorm:"column:cart_id;type:varchar(256);comment:购物车id;not null;default:[];" json:"cart_id"`                                               // 购物车id
	FreightPrice float64 `gorm:"column:freight_price;type:decimal(8, 2);comment:运费金额;not null;default:0.00;" json:"freight_price"`                                 // 运费金额
	TotalNum     uint32  `gorm:"column:total_num;type:int UNSIGNED;comment:订单商品总数;not null;default:0;" json:"total_num"`                                           // 订单商品总数
	TotalPrice   float64 `gorm:"column:total_price;type:decimal(8, 2) UNSIGNED;comment:订单总价;not null;" json:"total_price"`                                         // 订单总价
	TotalPostage float64 `gorm:"column:total_postage;type:decimal(8, 2) UNSIGNED;comment:邮费;not null;" json:"total_postage"`                                       // 邮费
	PayPrice     float64 `gorm:"column:pay_price;type:decimal(8, 2) UNSIGNED;comment:实际支付金额;not null;" json:"pay_price"`                                           // 实际支付金额
	PayPostage   float64 `gorm:"column:pay_postage;type:decimal(8, 2) UNSIGNED;comment:支付邮费;not null;" json:"pay_postage"`                                         // 支付邮费
	Paid         uint8   `gorm:"column:paid;type:tinyint UNSIGNED;comment:支付状态;not null;default:0;" json:"paid"`                                                   // 支付状态
	Status       int8    `gorm:"column:status;type:tinyint(1);comment:订单状态（-1 : 申请退款 -2 : 退货成功 0：待发货；1：待收货；2：已收货；3：待评价；-1：已退款）;not null;default:0;" json:"status"` // 订单状态（-1 : 申请退款 -2 : 退货成功 0：待发货；1：待收货；2：已收货；3：待评价；-1：已退款）
}

func (s *StoreOrder) TableName() string {
	return "store_order"
}

func (s *StoreOrder) CreateOrder() error {
	return global.DB.Create(&s).Error
}

func (s *StoreOrder) OrderList(name string) (ord []*StoreOrder, err error) {
	err = global.DB.Where("real_name = ?", name).Find(&ord).Error
	return ord, err
}
