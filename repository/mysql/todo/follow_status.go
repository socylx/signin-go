package todo

import (
	"time"
)

// FollowStatus 跟进状态表
type FollowStatus struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`               // 跟进状态名称
	CategoryID uint32    `gorm:"column:category_id" json:"category_id"` // 状态类别， 1售前(转介绍,新客户)，2续费
	Index      uint32    `gorm:"column:index" json:"index"`             // 显示排序
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *FollowStatus) TableName() string {
	return "follow_status"
}
