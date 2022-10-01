package todo

import (
	"time"
)

// LanguageMap [...]
type LanguageMap struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	LanguageID uint32    `gorm:"column:language_id" json:"language_id"`
	ForID      uint32    `gorm:"column:for_id" json:"for_id"`
	ForType    uint32    `gorm:"column:for_type" json:"for_type"` // 1-课程
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *LanguageMap) TableName() string {
	return "language_map"
}
