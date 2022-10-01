package todo

import (
	"time"
)

// Task [...]
type Task struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	Title        string    `gorm:"column:title" json:"title"`                   // 标题
	Content      string    `gorm:"column:content" json:"content"`               // todo内容
	SysMark      string    `gorm:"column:sys_mark" json:"sys_mark"`             // 系统备注
	Level        uint32    `gorm:"column:level" json:"level"`                   // 紧急程度，10：普通 ，20：重要，30：紧急
	DeadlineTime time.Time `gorm:"column:deadline_time" json:"deadline_time"`   // 截止时间
	TargetID     uint32    `gorm:"column:target_id" json:"target_id"`           // 目标Id，user_id / studio_id
	TargetType   uint32    `gorm:"column:target_type" json:"target_type"`       // 任务目标：1-学员, 2-门店, 3-公司
	TaskType     uint32    `gorm:"column:task_type" json:"task_type"`           // 任务类型： 1-排课, 2-激活, 3-等位失败
	Status       uint32    `gorm:"column:status" json:"status"`                 // 状态ID1:未完成，2:已完成, 3:结束
	CreateUserID uint32    `gorm:"column:create_user_id" json:"create_user_id"` // 创建者ID
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`       // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`       // 更新时间
	Remark       string    `gorm:"column:remark" json:"remark"`                 // 备注
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`                 // 是否删除
	CompleteID   string    `gorm:"column:complete_id" json:"complete_id"`       // 排课任务完成对应的 signin_id
	IsLong       bool      `gorm:"column:is_long" json:"is_long"`               // 是否是长期需求
}

// TableName get sql table name.获取数据库表名
func (m *Task) TableName() string {
	return "task"
}
