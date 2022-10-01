package todo

import (
	"time"
)

// TeacherApplyResources [...]
type TeacherApplyResources struct {
	ID             uint32    `gorm:"primaryKey;column:id" json:"id"`
	TeacherApplyID uint32    `gorm:"column:teacher_apply_id" json:"teacher_apply_id"`
	URL            string    `gorm:"column:url" json:"url"`
	Type           uint32    `gorm:"column:type" json:"type"` // 1、图片，2、视频，3、音频，4、其他
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel          bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TeacherApplyResources) TableName() string {
	return "teacher_apply_resources"
}
