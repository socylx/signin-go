package todo

import (
	"time"
)

// ShopAttributeValue 商品属性的值
type ShopAttributeValue struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"id"`
	ShopAttributeID uint32    `gorm:"column:shop_attribute_id" json:"shop_attribute_id"` // 所属的属性id
	Name            string    `gorm:"column:name" json:"name"`
	URL             string    `gorm:"column:url" json:"url"`   // 属性资源，例如颜色显示图片等
	Mark            string    `gorm:"column:mark" json:"mark"` // 备注
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopAttributeValue) TableName() string {
	return "shop_attribute_value"
}
