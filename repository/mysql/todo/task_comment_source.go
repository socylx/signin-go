package todo

import (
	"time"
)

// TaskCommentSource [...]
type TaskCommentSource struct {
	ID            uint32    `gorm:"primaryKey;column:id" json:"id"`
	TaskCommentID uint32    `gorm:"column:task_comment_id" json:"task_comment_id"`
	URL           string    `gorm:"column:url" json:"url"`
	SourceType    uint32    `gorm:"column:source_type" json:"source_type"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel         bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TaskCommentSource) TableName() string {
	return "task_comment_source"
}
