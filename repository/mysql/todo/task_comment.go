package todo

import (
	"time"
)

// TaskComment [...]
type TaskComment struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	PreID      uint32    `gorm:"column:pre_id" json:"pre_id"`
	TaskID     uint32    `gorm:"column:task_id" json:"task_id"`
	UserID     uint32    `gorm:"column:user_id" json:"user_id"`
	Content    string    `gorm:"column:content" json:"content"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TaskComment) TableName() string {
	return "task_comment"
}
