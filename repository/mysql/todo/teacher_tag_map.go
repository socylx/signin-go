package todo

import (
	"time"
)

// TeacherTagMap [...]
type TeacherTagMap struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	TeacherID    uint32    `gorm:"column:teacher_id" json:"teacher_id"`
	TeacherTagID uint32    `gorm:"column:teacher_tag_id" json:"teacher_tag_id"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TeacherTagMap) TableName() string {
	return "teacher_tag_map"
}
