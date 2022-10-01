package todo

import (
	"time"
)

// VideoRecord [...]
type VideoRecord struct {
	ID          uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID      uint32    `gorm:"column:user_id" json:"user_id"`
	ForID       uint32    `gorm:"column:for_id" json:"for_id"`
	ForType     uint32    `gorm:"column:for_type" json:"for_type"` // 1-activity_video, 2-录播课
	LastTime    float32   `gorm:"column:last_time" json:"last_time"`
	LongestTime float32   `gorm:"column:longest_time" json:"longest_time"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel       bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *VideoRecord) TableName() string {
	return "video_record"
}
