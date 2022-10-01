package todo

import (
	"time"
)

// VideoMark [...]
type VideoMark struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	ForID      uint32    `gorm:"column:for_id" json:"for_id"` // 对应的视频id
	ForType    uint32    `gorm:"column:for_type" json:"for_type"`
	StartTime  uint32    `gorm:"column:start_time" json:"start_time"` // 时间店，秒
	EndTime    uint32    `gorm:"column:end_time" json:"end_time"`
	Title      string    `gorm:"column:title" json:"title"` // 描述
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
}

// TableName get sql table name.获取数据库表名
func (m *VideoMark) TableName() string {
	return "video_mark"
}
