package todo

import (
	"time"
)

// Language [...]
type Language struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Key        string    `gorm:"column:key" json:"key"`
	Name       string    `gorm:"column:name" json:"name"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *Language) TableName() string {
	return "language"
}
