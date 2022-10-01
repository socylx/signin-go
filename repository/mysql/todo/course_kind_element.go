package todo

import (
	"time"
)

// CourseKindElement [...]
type CourseKindElement struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	CourseKindID uint32    `gorm:"column:course_kind_id" json:"course_kind_id"` // 对应的课程id
	Title        string    `gorm:"column:title" json:"title"`                   // 元素名字
	Decs         string    `gorm:"column:decs" json:"decs"`                     // 视频描述
	ImgURL       string    `gorm:"column:img_url" json:"img_url"`               // 介绍的图片链接
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *CourseKindElement) TableName() string {
	return "course_kind_element"
}
