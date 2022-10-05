package judge_user_log

import (
	"time"
)

// JudgeUserLog [...]
type JudgeUserLog struct {
	ID          uint32    `gorm:"primaryKey;column:id" json:"id"`
	JudgeUserID uint32    `gorm:"column:judge_user_id" json:"judge_user_id"`
	Status      uint32    `gorm:"column:status" json:"status"`           // 本次状态
	Mark        string    `gorm:"column:mark" json:"mark"`               // 操作备注
	SysMark     string    `gorm:"column:sys_mark" json:"sys_mark"`       // 系统备注
	OptUserID   uint32    `gorm:"column:opt_user_id" json:"opt_user_id"` // 操作人
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel       bool      `gorm:"column:is_del" json:"is_del"`
}
