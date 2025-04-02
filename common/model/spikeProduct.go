package model

import (
	"gorm.io/gorm"
	"miaosha-jjl/common/global"
)

type SpikeProduct struct {
	gorm.Model
	ProductId    int32   `gorm:"column:product_id;type:int;comment:商品id;not null;" json:"product_id"`                  // 商品id
	ProductName  string  `gorm:"column:product_name;type:varchar(30);comment:商品名称;not null;" json:"product_name"`      // 商品名称
	ProductPrice float64 `gorm:"column:product_price;type:decimal(10, 2);comment:商品原价;not null;" json:"product_price"` // 商品原价
	SpikePrice   float64 `gorm:"column:spike_price;type:decimal(10, 2);comment:秒杀价;not null;" json:"spike_price"`      // 秒杀价
	SpikeNum     int32   `gorm:"column:spike_num;type:int;comment:秒杀库存;not null;" json:"spike_num"`                    // 秒杀库存
	StartTime    string  `gorm:"column:start_time;type:varchar(100);comment:秒杀开始时间;not null;" json:"start_time"`       // 秒杀开始时间
	EndTime      string  `gorm:"column:end_time;type:varchar(100);comment:秒杀结束时间;not null;" json:"end_time"`           // 秒杀结束时间
	SpikeStatus  int32   `gorm:"column:spike_status;type:int;comment:1准备中2进行中3收尾中;not null;" json:"spike_status"`      // 1准备中2进行中3收尾中
}

func (p *SpikeProduct) TableName() string {
	return "spike_product"
}

func (s *SpikeProduct) AddSpikeProduct() error {
	return global.DB.Create(&s).Error
}

func (s *SpikeProduct) QuerySpikeProduct(id int) (product *SpikeProduct, err error) {
	err = global.DB.Where("id = ?", id).Limit(1).Find(&s).Error
	return product, err
}
