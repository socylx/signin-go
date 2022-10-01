package todo

import (
	"time"
)

// TaskContent [...]
type TaskContent struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	TaskID     uint32    `gorm:"column:task_id" json:"task_id"`         // 任务Id
	Value      string    `gorm:"column:value" json:"value"`             // 意向值
	ValueType  uint32    `gorm:"column:value_type" json:"value_type"`   // 意向类型, 1-weekday_time,2-course_kind,3-course_grade,4-course_level,5-teacher,6-studio
	IsKey      bool      `gorm:"column:is_key" json:"is_key"`           // 是否参与约课信息校验
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`           // 是否删除
}

// TableName get sql table name.获取数据库表名
func (m *TaskContent) TableName() string {
	return "task_content"
}
