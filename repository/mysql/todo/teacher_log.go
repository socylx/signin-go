package todo

import (
	"time"
)

// TeacherLog [...]
type TeacherLog struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	TeacherID  uint32    `gorm:"column:teacher_id" json:"teacher_id"`
	LogContent string    `gorm:"column:log_content" json:"log_content"`
	SysRemark  string    `gorm:"column:sys_remark" json:"sys_remark"`
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TeacherLog) TableName() string {
	return "teacher_log"
}
