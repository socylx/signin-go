package todo

import (
	"time"
)

// VideoCollect 用户视频收藏表
type VideoCollect struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID          uint32    `gorm:"column:user_id" json:"user_id"`                     // 属于哪个用户的收藏
	ActivityVideoID uint32    `gorm:"column:activity_video_id" json:"activity_video_id"` // 收藏的视频id
	ActivityID      uint32    `gorm:"column:activity_id" json:"activity_id"`
	Title           string    `gorm:"column:title" json:"title"` // 标题，存视频名称
	Type            uint32    `gorm:"column:type" json:"type"`   // 1课程视频，2结课视频
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *VideoCollect) TableName() string {
	return "video_collect"
}
