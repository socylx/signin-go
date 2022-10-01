package todo

import (
	"time"
)

// Resource [...]
type Resource struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	ForID        uint32    `gorm:"column:for_id" json:"for_id"`
	ForType      uint32    `gorm:"column:for_type" json:"for_type"`
	ResourceURL  string    `gorm:"column:resource_url" json:"resource_url"`
	ResourceType uint32    `gorm:"column:resource_type" json:"resource_type"`
	ResourceName string    `gorm:"column:resource_name" json:"resource_name"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
	OptUserID    uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
}

// TableName get sql table name.获取数据库表名
func (m *Resource) TableName() string {
	return "resource"
}
