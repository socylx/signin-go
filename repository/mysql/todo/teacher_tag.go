package todo

import (
	"time"
)

// TeacherTag [...]
type TeacherTag struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Type       uint32    `gorm:"column:type" json:"type"` // 1-教师标签, 2-教师申请标签
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TeacherTag) TableName() string {
	return "teacher_tag"
}
