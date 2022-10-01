package todo

import (
	"time"
)

// ShopSkuCategory [...]
type ShopSkuCategory struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"` // 分类名
	Logo       string    `gorm:"column:logo" json:"logo"`
	Index      uint32    `gorm:"column:index" json:"index"`     // 排序, 越小越靠前
	IsShow     bool      `gorm:"column:is_show" json:"is_show"` // 是否展示
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopSkuCategory) TableName() string {
	return "shop_sku_category"
}
