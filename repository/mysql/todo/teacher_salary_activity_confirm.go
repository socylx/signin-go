package todo

import (
	"time"
)

// TeacherSalaryActivityConfirm [...]
type TeacherSalaryActivityConfirm struct {
	ID                      uint32    `gorm:"primaryKey;column:id" json:"id"`
	TeacherSalaryActivityID uint32    `gorm:"column:teacher_salary_activity_id" json:"teacher_salary_activity_id"` // 课程账单ID
	SysMark                 string    `gorm:"column:sys_mark" json:"sys_mark"`                                     // 系统日志
	UserID                  uint32    `gorm:"column:user_id" json:"user_id"`                                       // 账单确认的user_id
	CreateTime              time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime              time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel                   bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TeacherSalaryActivityConfirm) TableName() string {
	return "teacher_salary_activity_confirm"
}
