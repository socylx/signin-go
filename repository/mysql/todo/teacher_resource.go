package todo

import (
	"time"
)

// TeacherResource [...]
type TeacherResource struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	TeacherID    uint32    `gorm:"column:teacher_id" json:"teacher_id"`
	URL          string    `gorm:"column:url" json:"url"`
	ResourceType uint32    `gorm:"column:resource_type" json:"resource_type"` // 资源类型. 1-图片,2-视频...
	Type         uint32    `gorm:"column:type" json:"type"`                   // 资源类别. 1-教师身份证图片
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TeacherResource) TableName() string {
	return "teacher_resource"
}
