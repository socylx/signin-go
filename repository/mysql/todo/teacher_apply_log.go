package todo

import (
	"time"
)

// TeacherApplyLog [...]
type TeacherApplyLog struct {
	ID             uint32    `gorm:"primaryKey;column:id" json:"id"`
	TeacherApplyID uint32    `gorm:"column:teacher_apply_id" json:"teacher_apply_id"`
	LogContent     string    `gorm:"column:log_content" json:"log_content"`
	SysRemark      string    `gorm:"column:sys_remark" json:"sys_remark"`
	OptUserID      uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel          bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TeacherApplyLog) TableName() string {
	return "teacher_apply_log"
}
