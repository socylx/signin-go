package todo

import (
	"time"
)

// RolePage [...]
type RolePage struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	SystemPageID uint32    `gorm:"column:system_page_id" json:"system_page_id"` // system_page 表id，
	RoleID       uint32    `gorm:"column:role_id" json:"role_id"`
	Name         string    `gorm:"column:name" json:"name"`
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *RolePage) TableName() string {
	return "role_page"
}
