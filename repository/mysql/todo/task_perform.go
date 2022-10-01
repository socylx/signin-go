package todo

import (
	"time"
)

// TaskPerform [...]
type TaskPerform struct {
	ID          uint32    `gorm:"primaryKey;column:id" json:"id"`
	TaskID      uint32    `gorm:"column:task_id" json:"task_id"`           // 任务ID
	PerformType uint32    `gorm:"column:perform_type" json:"perform_type"` // 1-个人执行，2-门店集体执行
	PerformID   uint32    `gorm:"column:perform_id" json:"perform_id"`     // user_id或studio_id
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`   // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`   // 更新时间
	IsDel       bool      `gorm:"column:is_del" json:"is_del"`             // 是否删除
}

// TableName get sql table name.获取数据库表名
func (m *TaskPerform) TableName() string {
	return "task_perform"
}
