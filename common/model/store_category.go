package model

import (
	"miaosha-jjl/common/global"
)

type StoreCategory struct {
	ID       int32           `gorm:"column:id;primaryKey;autoIncrement" json:"id"`                 // 商品分类表ID
	Pid      int32           `gorm:"column:pid;not null" json:"pid"`                               // 父id
	CateName string          `gorm:"column:cate_name;type:varchar(100);not null" json:"cate_name"` // 分类名称
	Sort     int32           `gorm:"column:sort;not null" json:"sort"`                             // 排序
	Pic      string          `gorm:"column:pic;type:varchar(128);default:''" json:"pic"`           // 图标
	IsShow   int8            `gorm:"column:is_show;type:tinyint(1);default:1" json:"is_show"`      // 是否推荐
	AddTime  int32           `gorm:"column:add_time;not null" json:"add_time"`                     // 添加时间
	Children []StoreCategory `json:"children" gorm:"-"`
}

// TableName 设置表名
func (s *StoreCategory) TableName() string {
	return "store_category"
}

// 获取指定父分类的所有子分类
func (s *StoreCategory) GetChild(pid int) ([]StoreCategory, error) {
	var cate []StoreCategory
	err := global.DB.Where("pid = ?", pid).Find(&cate).Error
	return cate, err
}
func (s *StoreCategory) GetCateWithChild(pid int) (StoreCategory, error) {
	var cate StoreCategory
	// 获取当前分类
	err := global.DB.Where("id = ?", pid).First(&cate).Error
	if err != nil {
		return StoreCategory{}, err
	}

	// 获取当前分类的所有子分类
	child, err := s.GetChild(pid)
	if err != nil {
		return StoreCategory{}, err
	}

	// 将子分类设置到当前分类的 Children 字段
	cate.Children = child

	// 递归处理每一个子分类
	for i := range cate.Children {
		withChild, err := s.GetCateWithChild(int(cate.Children[i].ID))
		if err != nil {
			continue // 如果某个子分类的递归失败，跳过
		}
		cate.Children[i] = withChild
	}

	return cate, nil
}
