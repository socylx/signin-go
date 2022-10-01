package todo

import (
	"time"
)

// UserTipsLog 用户通知状态表，用来标记某类通知对用户已经发送过了，不用在发送
type UserTipsLog struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID     uint32    `gorm:"column:user_id" json:"user_id"`   // 用户id
	TipsKey    string    `gorm:"column:tips_key" json:"tips_key"` // 提示的关键字，原来标记是什么提示
	Status     uint32    `gorm:"column:status" json:"status"`     // 0未发送，1已经发送，2手动关闭，3程序关闭
	Deadline   time.Time `gorm:"column:deadline" json:"deadline"` // 下次允许的通知时间；例如某类通知是一个月之后可以再次发送
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *UserTipsLog) TableName() string {
	return "user_tips_log"
}
