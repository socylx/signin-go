package todo

import (
	"time"
)

// Subscribe 订阅
type Subscribe struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID     uint32    `gorm:"column:user_id" json:"user_id"`         // 订阅所属人id
	Type       uint32    `gorm:"column:type" json:"type"`               // 订阅类型，1课程
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"` // 操作人id
	Mark       string    `gorm:"column:mark" json:"mark"`               // 备注
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *Subscribe) TableName() string {
	return "subscribe"
}
