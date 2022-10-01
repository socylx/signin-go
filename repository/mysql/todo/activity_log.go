package todo

import (
	"time"
)

// ActivityLog [...]
type ActivityLog struct {
	ID            uint32    `gorm:"primaryKey;column:id" json:"id"`
	PrevTeacherID uint32    `gorm:"column:prev_teacher_id" json:"prev_teacher_id"` // 上一个上课老师id，如果没变更过，则和teacher_id值相等
	ActivityID    uint32    `gorm:"column:activity_id" json:"activity_id"`         // 修改日志对应的课程ID
	CourseID      uint32    `gorm:"column:course_id" json:"course_id"`             // 修改日志对应的课程组ID
	OptUserID     uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`         // 修改人ID
	LogContent    string    `gorm:"column:log_content" json:"log_content"`         // 修改日志
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel         bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ActivityLog) TableName() string {
	return "activity_log"
}
