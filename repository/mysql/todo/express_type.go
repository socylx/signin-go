package todo

import (
	"time"
)

// ExpressType [...]
type ExpressType struct {
	ID               uint32    `gorm:"primaryKey;column:id" json:"id"`
	ExpressCompanyID uint32    `gorm:"column:express_company_id" json:"express_company_id"`
	Code             string    `gorm:"column:code" json:"code"`
	Type             string    `gorm:"column:type" json:"type"`
	Desc             string    `gorm:"column:desc" json:"desc"`
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel            bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ExpressType) TableName() string {
	return "express_type"
}
