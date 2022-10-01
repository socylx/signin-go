package todo

import (
	"time"
)

// Like 存储小程序生成的 formId, 用来推送
type Like struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Type       uint32    `gorm:"column:type" json:"type"`       // 1、发现页视频，2课堂视频
	ForID      uint32    `gorm:"column:for_id" json:"for_id"`   // 为什么点赞
	UserID     uint32    `gorm:"column:user_id" json:"user_id"` // 用户id
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *Like) TableName() string {
	return "like"
}
