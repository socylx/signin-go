package todo

import (
	"time"
)

// VideoLog [...]
type VideoLog struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	ForID      uint32    `gorm:"column:for_id" json:"for_id"` // 对应的视频id
	ForType    uint32    `gorm:"column:for_type" json:"for_type"`
	LogContent string    `gorm:"column:log_content" json:"log_content"` // 描述
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *VideoLog) TableName() string {
	return "video_log"
}
