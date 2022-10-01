package todo

import (
	"time"
)

// ActivityTeacher [...]
type ActivityTeacher struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"id"`
	ActivityID      uint32    `gorm:"column:activity_id" json:"activity_id"`             // 对应的课程id
	TeacherID       uint32    `gorm:"column:teacher_id" json:"teacher_id"`               // 教课老师id
	TeacherComeTime time.Time `gorm:"column:teacher_come_time" json:"teacher_come_time"` // 老师到店时间
	OptType         uint32    `gorm:"column:opt_type" json:"opt_type"`                   // 操作类型，1正常，2代课，
	PreID           uint32    `gorm:"column:pre_id" json:"pre_id"`
	InClass         bool      `gorm:"column:in_class" json:"in_class"` // 是否上课，1是，0否
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`
	OptUserID       uint32    `gorm:"column:opt_user_id" json:"opt_user_id"` // 操作人
}

// TableName get sql table name.获取数据库表名
func (m *ActivityTeacher) TableName() string {
	return "activity_teacher"
}
