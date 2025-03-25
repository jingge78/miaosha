package model

import "miaosha-jjl/common/global"

type Product struct {
	Id        uint64  `gorm:"column:id;type:mediumint;comment:商品id;primaryKey;not null;" json:"id"`                                     // 商品id
	MerId     uint64  `gorm:"column:mer_id;type:int UNSIGNED;comment:商户Id(0为总后台管理员创建,不为0的时候是商户后台创建);not null;default:0;" json:"mer_id"` // 商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)
	Image     string  `gorm:"column:image;type:varchar(256);comment:商品图片;not null;" json:"image"`                                       // 商品图片
	StoreName string  `gorm:"column:store_name;type:varchar(128);comment:商品名称;not null;" json:"store_name"`                             // 商品名称
	StoreInfo string  `gorm:"column:store_info;type:varchar(256);comment:商品简介;not null;" json:"store_info"`                             // 商品简介
	Price     float64 `gorm:"column:price;type:decimal(8, 2) UNSIGNED;comment:商品价格;not null;default:0.00;" json:"price"`                // 商品价格
	VipPrice  float64 `gorm:"column:vip_price;type:decimal(8, 2) UNSIGNED;comment:会员价格;not null;default:0.00;" json:"vip_price"`        // 会员价格
	Postage   float64 `gorm:"column:postage;type:decimal(8, 2) UNSIGNED;comment:邮费;not null;default:0.00;" json:"postage"`              // 邮费
	Stock     uint64  `gorm:"column:stock;type:mediumint UNSIGNED;comment:库存;not null;default:0;" json:"stock"`                         // 库存
	IsShow    uint64  `gorm:"column:is_show;type:tinyint(1);comment:状态（0：未上架，1：上架）;not null;default:1;" json:"is_show"`                 // 状态（0：未上架，1：上架）
	AddTime   int     `gorm:"column:add_time;type:int UNSIGNED;comment:添加时间;not null;" json:"add_time"`                                 // 添加时间
	IsPostage uint64  `gorm:"column:is_postage;type:tinyint UNSIGNED;comment:是否包邮;not null;default:0;" json:"is_postage"`               // 是否包邮
	IsDel     uint8   `gorm:"column:is_del;type:tinyint UNSIGNED;comment:是否删除;not null;default:0;" json:"is_del"`                       // 是否删除
	IsSeckill uint64  `gorm:"column:is_seckill;type:tinyint UNSIGNED;comment:秒杀状态 0 未开启 1已开启;not null;default:0;" json:"is_seckill"`    // 秒杀状态 0 未开启 1已开启
	Browse    uint64  `gorm:"column:browse;type:int;comment:浏览量;default:0;" json:"browse"`                                              // 浏览量
}

<<<<<<< HEAD
func (p *Product) TableName() string {
	return "product"
}
func (p *Product) ProductDetailReq(productId uint64) error {
	return global.DB.Where("id = ?", productId).Limit(1).Find(&p).Error
}
func (p *Product) GetAllProduct() ([]Product, error) {
	var allProduct []Product
	err := global.DB.Find(&allProduct).Error
	return allProduct, err
=======
func (p *Product) ProductDetailReq(productId uint64) error {
	return global.DB.Debug().Table("product").Where("id = ?", productId).Limit(1).Find(&p).Error
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
}
