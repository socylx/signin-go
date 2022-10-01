package todo

import (
	"time"
)

// TaskSource [...]
type TaskSource struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	TaskID     uint32    `gorm:"column:task_id" json:"task_id"`         // 任务ID
	URL        string    `gorm:"column:url" json:"url"`                 // 资源链接
	SourceType uint32    `gorm:"column:source_type" json:"source_type"` // 资源类型,1:图片,2:视频，3:音频
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`           // 是否删除
}

// TableName get sql table name.获取数据库表名
func (m *TaskSource) TableName() string {
	return "task_source"
}
