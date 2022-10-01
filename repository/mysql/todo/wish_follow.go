package todo

import (
	"time"
)

// WishFollow [...]
type WishFollow struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	WishID     uint32    `gorm:"column:wish_id" json:"wish_id"`
	UserID     uint32    `gorm:"column:user_id" json:"user_id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *WishFollow) TableName() string {
	return "wish_follow"
}
