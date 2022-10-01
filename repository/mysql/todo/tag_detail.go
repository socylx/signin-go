package todo

import (
	"time"
)

// TagDetail [...]
type TagDetail struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	TagID      uint32    `gorm:"column:tag_id" json:"tag_id"`
	Name       string    `gorm:"column:name" json:"name"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TagDetail) TableName() string {
	return "tag_detail"
}
