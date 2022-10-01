package todo

import (
	"time"
)

// OnlineVideoRecord [...]
type OnlineVideoRecord struct {
	ID            uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID        uint32    `gorm:"column:user_id" json:"user_id"`
	OnlineVideoID uint32    `gorm:"column:online_video_id" json:"online_video_id"`
	LastTime      float32   `gorm:"column:last_time" json:"last_time"`
	LongestTime   float32   `gorm:"column:longest_time" json:"longest_time"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel         bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *OnlineVideoRecord) TableName() string {
	return "online_video_record"
}
