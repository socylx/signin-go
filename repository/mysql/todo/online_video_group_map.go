package todo

import (
	"time"
)

// OnlineVideoGroupMap [...]
type OnlineVideoGroupMap struct {
	OnlineVideoID      uint32    `gorm:"primaryKey;column:online_video_id" json:"online_video_id"`             // 对应的录播视频id
	OnlineVideoGroupID uint32    `gorm:"primaryKey;column:online_video_group_id" json:"online_video_group_id"` // 对应的录播课组id
	Index              int       `gorm:"column:index" json:"index"`                                            // 排序
	UpdateTime         time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime         time.Time `gorm:"column:create_time" json:"create_time"`
	IsDel              bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *OnlineVideoGroupMap) TableName() string {
	return "online_video_group_map"
}
