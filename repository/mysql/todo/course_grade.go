package todo

import (
	"time"
)

// CourseGrade 课程等级
type CourseGrade struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`     // 名称
	IsDel      bool      `gorm:"column:is_del" json:"is_del"` // 是否被删除
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	StartNum   uint32    `gorm:"column:start_num" json:"start_num"` // 星级难度，满分 10
}

// TableName get sql table name.获取数据库表名
func (m *CourseGrade) TableName() string {
	return "course_grade"
}
