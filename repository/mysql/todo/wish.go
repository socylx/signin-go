package todo

import (
	"time"
)

// Wish [...]
type Wish struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID     uint32    `gorm:"column:user_id" json:"user_id"`
	Type       uint32    `gorm:"column:type" json:"type"`
	Name       string    `gorm:"column:name" json:"name"`
	URL        string    `gorm:"column:url" json:"url"`
	Desc       string    `gorm:"column:desc" json:"desc"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	Status     uint32    `gorm:"column:status" json:"status"`
}

// TableName get sql table name.获取数据库表名
func (m *Wish) TableName() string {
	return "wish"
}
