package todo

import (
	"time"
)

// OnlineVideoTagMap [...]
type OnlineVideoTagMap struct {
	OnlineVideoID    uint32    `gorm:"primaryKey;column:online_video_id" json:"online_video_id"`         // 视频id
	OnlineVideoTagID uint32    `gorm:"primaryKey;column:online_video_tag_id" json:"online_video_tag_id"` // 视频标签id
	UpdateTime       time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time"`
	IsDel            bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *OnlineVideoTagMap) TableName() string {
	return "online_video_tag_map"
}
