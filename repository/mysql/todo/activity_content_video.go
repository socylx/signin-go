package todo

import (
	"time"
)

// ActivityContentVideo [...]
type ActivityContentVideo struct {
	ID            uint32    `gorm:"primaryKey;column:id" json:"id"`
	ActivityID    uint32    `gorm:"column:activity_id" json:"activity_id"`         // 对应的课程id
	OnlineVideoID uint32    `gorm:"column:online_video_id" json:"online_video_id"` // 对应的视频
	Index         int       `gorm:"column:index" json:"index"`                     // 排序
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	IsDel         bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ActivityContentVideo) TableName() string {
	return "activity_content_video"
}
