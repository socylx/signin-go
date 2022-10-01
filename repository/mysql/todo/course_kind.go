package todo

import (
	"time"
)

// CourseKind [...]
type CourseKind struct {
	ID          uint32    `gorm:"primaryKey;column:id" json:"id"`
	Type        string    `gorm:"column:type" json:"type"`                 // 舞种名称
	Desc        string    `gorm:"column:desc" json:"desc"`                 // 舞种描述
	DescPopular string    `gorm:"column:desc_popular" json:"desc_popular"` // 通俗版舞种介绍
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`   // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`   // 更新时间
	IsDel       bool      `gorm:"column:is_del" json:"is_del"`             // 1代表删除
}

// TableName get sql table name.获取数据库表名
func (m *CourseKind) TableName() string {
	return "course_kind"
}
