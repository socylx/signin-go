package todo

import (
	"time"
)

// ResourceAttribute [...]
type ResourceAttribute struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	ForID      uint32    `gorm:"column:for_id" json:"for_id"`
	ForType    uint32    `gorm:"column:for_type" json:"for_type"`
	Value      uint32    `gorm:"column:value" json:"value"`
	ValueType  uint32    `gorm:"column:value_type" json:"value_type"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ResourceAttribute) TableName() string {
	return "resource_attribute"
}
