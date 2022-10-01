package todo

import (
	"time"
)

// ShopCategory 商品分类
type ShopCategory struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	ParentID   uint32    `gorm:"column:parent_id" json:"parent_id"` // 付级分类id, 如果为0则代表第一层级
	URL        string    `gorm:"column:url" json:"url"`             // 资源
	Name       string    `gorm:"column:name" json:"name"`           // 名称
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopCategory) TableName() string {
	return "shop_category"
}
