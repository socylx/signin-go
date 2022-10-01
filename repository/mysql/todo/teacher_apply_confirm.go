package todo

import (
	"time"
)

// TeacherApplyConfirm [...]
type TeacherApplyConfirm struct {
	ID             uint32    `gorm:"primaryKey;column:id" json:"id"`
	TeacherApplyID uint32    `gorm:"column:teacher_apply_id" json:"teacher_apply_id"`
	Status         uint32    `gorm:"column:status" json:"status"` // 状态
	Content        string    `gorm:"column:content" json:"content"`
	Amount         uint32    `gorm:"column:amount" json:"amount"`
	Type           uint32    `gorm:"column:type" json:"type"` // 1-师资确认记录, 2-全职老师确认记录
	OptUserID      uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel          bool      `gorm:"column:is_del" json:"is_del"`
	BelongID       uint32    `gorm:"column:belong_id" json:"belong_id"` // 谁的确认
}

// TableName get sql table name.获取数据库表名
func (m *TeacherApplyConfirm) TableName() string {
	return "teacher_apply_confirm"
}
