package todo

import (
	"time"
)

// ExpressCompany [...]
type ExpressCompany struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Code       string    `gorm:"column:code" json:"code"`
	Logo       string    `gorm:"column:logo" json:"logo"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ExpressCompany) TableName() string {
	return "express_company"
}
