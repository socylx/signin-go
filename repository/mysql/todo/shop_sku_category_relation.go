package todo

import (
	"time"
)

// ShopSkuCategoryRelation [...]
type ShopSkuCategoryRelation struct {
	ID            uint32    `gorm:"primaryKey;column:id" json:"id"`
	SkuCategoryID uint32    `gorm:"column:sku_category_id" json:"sku_category_id"`
	SkuID         uint32    `gorm:"column:sku_id" json:"sku_id"`
	Index         uint32    `gorm:"column:index" json:"index"` // 排序
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel         bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopSkuCategoryRelation) TableName() string {
	return "shop_sku_category_relation"
}
