package todo

import (
	"time"
)

// Tag [...]
type Tag struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Type       uint32    `gorm:"column:type" json:"type"` // 1-线索标签,2-系列课标签
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *Tag) TableName() string {
	return "tag"
}
