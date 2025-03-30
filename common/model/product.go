package model

import (
	"miaosha-jjl/common/global"
	"miaosha-jjl/common/utils"
)

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
}

// 价格查找//jj
func (p *Product) PriceFind(price float64) (Pro []*Product, err error) {
	err = global.DB.Where("price = ?", price).Find(&Pro).Error
	return Pro, err
}

// 按价格和上架状态排序
func (p *Product) ProductSortByIsShowOrPrice(isShow int) ([]Product, error) {

	//商品价格排序
	//商品状态  0 下架 1 上架
	//原型商品可以根据销量，价格，评价，和最新上架来进行商品排序

	var product []Product
	err := global.DB.Order("price DESC").Where("is_show = ?", isShow).Find(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) ProductFilter(minPrice, maxPrice float32, name string, isPostage, page, pageSize int64) (product []Product, pagination utils.Pagination, err error) {

	var products []Product

	query := global.DB.Model(&Product{}).Where("is_del = ?", 0) //是否删除

	//筛选条件
	if minPrice > 0 {
		query = query.Where("price >=?", minPrice)
	}
	if maxPrice > 0 {
		query = query.Where("price <=?", maxPrice)
	}
	if name != "" {
		query = query.Where("store_name LIKE ?", "%"+name+"%")
	}
	if isPostage >= 0 {
		query = query.Where("is_postage = ?", isPostage)
	}

	pagination = utils.Pagination{
		Page:     int(page),
		PageSize: int(pageSize),
	}
	pagination.Validate() // 验证分页参数
	// 先获取总数
	var total int64
	if err = query.Count(&total).Error; err != nil {
		return nil, pagination, err
	}
	pagination.SetTotal(total)
	err = query.Offset(pagination.Offset()).Limit(pagination.Limit()).Find(&products).Error
	if err != nil {
		return nil, pagination, err
	}

	return products, pagination, nil
}
