package todo

import (
	"time"
)

// OnlineVideoKindElement [...]
type OnlineVideoKindElement struct {
	ID                  uint32    `gorm:"primaryKey;column:id" json:"id"`
	OnlineVideoID       uint32    `gorm:"column:online_video_id" json:"online_video_id"`               // 对应的视频id
	CourseKindElementID uint32    `gorm:"column:course_kind_element_id" json:"course_kind_element_id"` // 舞种元素id
	UpdateTime          time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime          time.Time `gorm:"column:create_time" json:"create_time"`
	IsDel               bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *OnlineVideoKindElement) TableName() string {
	return "online_video_kind_element"
}
