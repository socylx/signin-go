package todo

import (
	"time"
)

// TeacherSalaryActivityLog [...]
type TeacherSalaryActivityLog struct {
	ID                      uint32    `gorm:"primaryKey;column:id" json:"id"`
	TeacherSalaryActivityID uint32    `gorm:"column:teacher_salary_activity_id" json:"teacher_salary_activity_id"` // 结算表 ID
	LogContent              string    `gorm:"column:log_content" json:"log_content"`                               // 结算操作日志
	OptUserID               uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`                               // 操作人
	CreateTime              time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime              time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel                   bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TeacherSalaryActivityLog) TableName() string {
	return "teacher_salary_activity_log"
}
