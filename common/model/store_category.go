package model

type StoreCategory struct {
	Id       int32  `gorm:"column:id;type:mediumint;comment:商品分类表ID;primaryKey;not null;" json:"id"`        // 商品分类表ID
	Pid      int32  `gorm:"column:pid;type:mediumint;comment:父id;not null;" json:"pid"`                     // 父id
	CateName string `gorm:"column:cate_name;type:varchar(100);comment:分类名称;not null;" json:"cate_name"`     // 分类名称
	Sort     int32  `gorm:"column:sort;type:mediumint;comment:排序;not null;" json:"sort"`                    // 排序
	Pic      string `gorm:"column:pic;type:varchar(128);comment:图标;not null;" json:"pic"`                   // 图标
	IsShow   int8   `gorm:"column:is_show;type:tinyint(1);comment:是否推荐;not null;default:1;" json:"is_show"` // 是否推荐
	AddTime  int32  `gorm:"column:add_time;type:int;comment:添加时间;not null;" json:"add_time"`                // 添加时间
}
