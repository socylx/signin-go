package todo

import (
	"time"
)

// OnlineVideoTag [...]
type OnlineVideoTag struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"` // 名称
	Decs       string    `gorm:"column:decs" json:"decs"` // 视频描述
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *OnlineVideoTag) TableName() string {
	return "online_video_tag"
}
