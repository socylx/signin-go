package todo

import (
	"time"
)

// WishResult [...]
type WishResult struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	WishID     uint32    `gorm:"column:wish_id" json:"wish_id"`
	ForID      uint32    `gorm:"column:for_id" json:"for_id"`
	ForType    uint32    `gorm:"column:for_type" json:"for_type"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *WishResult) TableName() string {
	return "wish_result"
}
