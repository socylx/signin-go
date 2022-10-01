package todo

import (
	"time"
)

// Source 用户来源表
type Source struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"` // 来源名称
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	Index      uint32    `gorm:"column:index" json:"index"` // 排序
}

// TableName get sql table name.获取数据库表名
func (m *Source) TableName() string {
	return "source"
}
