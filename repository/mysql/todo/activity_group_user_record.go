package todo

import (
	"time"
)

// ActivityGroupUserRecord [...]
type ActivityGroupUserRecord struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"id"`
	ActivityGroupID uint32    `gorm:"column:activity_group_id" json:"activity_group_id"`
	UserID          uint32    `gorm:"column:user_id" json:"user_id"`
	OrderID         uint32    `gorm:"column:order_id" json:"order_id"`
	Status          uint32    `gorm:"column:status" json:"status"` // 1-正常,2-终止
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ActivityGroupUserRecord) TableName() string {
	return "activity_group_user_record"
}
