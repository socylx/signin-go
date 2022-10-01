package todo

import (
	"time"
)

// Brand [...]
type Brand struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`               // 品牌名称
	LogoURL    string    `gorm:"column:logo_url" json:"logo_url"`       // logo web  url地址
	Desc       string    `gorm:"column:desc" json:"desc"`               // 品牌描述
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`           // 是否删除
}

// TableName get sql table name.获取数据库表名
func (m *Brand) TableName() string {
	return "brand"
}
