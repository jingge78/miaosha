package model

import (
	"miaosha-jjl/common/global"
	"time"
)

type EbStoreProductRelation struct {
	Uid       int       `gorm:"column:uid" json:"uid"`               //用户ID
	ProductId int       `gorm:"column:product_id" json:"product_id"` //商品id
	Type      string    `gorm:"column:type" json:"type"`             //类型(收藏(collect）、点赞(like))
	Category  string    `gorm:"column:category" json:"category"`     //某种类型的商品(普通商品、秒杀商品)
	AddTime   time.Time `gorm:"column:add_time" json:"add_time"`     //添加时间
}

func (e *EbStoreProductRelation) TableName() string {
	return "eb_store_product_relation"
}
func (e *EbStoreProductRelation) Create() error {
	return global.DB.Create(&e).Error
}
func (e *EbStoreProductRelation) GetCollectProduct(uid int) ([]EbStoreProductRelation, error) {
	var product []EbStoreProductRelation
	err := global.DB.Where("uid = ?", uid).Where("type = ?", "收藏").Find(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
