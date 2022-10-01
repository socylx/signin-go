package todo

import (
	"time"
)

// TeacherFollow [...]
type TeacherFollow struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	ForID      uint32    `gorm:"column:for_id" json:"for_id"`
	ForType    uint32    `gorm:"column:for_type" json:"for_type"` // 1-对简历的跟进，2-对老师的跟进
	ActivityID uint32    `gorm:"column:activity_id" json:"activity_id"`
	Content    string    `gorm:"column:content" json:"content"`
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TeacherFollow) TableName() string {
	return "teacher_follow"
}
