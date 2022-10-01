package todo

import (
	"time"
)

// Permission [...]
type Permission struct {
	ID            uint32    `gorm:"primaryKey;column:id" json:"id"`
	Type          int       `gorm:"column:type" json:"type"`                       // 权限类型:1
	Desc          string    `gorm:"column:desc" json:"desc"`                       // 描述
	ApplyStudioID int       `gorm:"column:apply_studio_id" json:"apply_studio_id"` // 权限适用工作室场馆
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`         // 更新时间
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`         // 创建时间
	IsDel         int8      `gorm:"column:is_del" json:"is_del"`                   // 是否删除
	RoleID        int       `gorm:"column:role_id" json:"role_id"`
}

// TableName get sql table name.获取数据库表名
func (m *Permission) TableName() string {
	return "permission"
}
