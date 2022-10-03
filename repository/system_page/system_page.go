package system_page

import (
	"time"
)

// SystemPage 后端自定义的权限组
type SystemPage struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	PageKey    string    `gorm:"column:page_key" json:"page_key"`
	Name       string    `gorm:"column:name" json:"name"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func tableName() string {
	return "system_page"
}
