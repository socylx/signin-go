package todo

import (
	"time"
)

// WxMsg [...]
type WxMsg struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	MsgID        string    `gorm:"column:msg_id" json:"msg_id"`                 // 消息ID
	StaffUserID  uint32    `gorm:"column:staff_user_id" json:"staff_user_id"`   // 工作人员user_id
	FriendUserID uint32    `gorm:"column:friend_user_id" json:"friend_user_id"` // 学员user_id
	UserBeforeID uint32    `gorm:"column:user_before_id" json:"user_before_id"`
	Content      string    `gorm:"column:content" json:"content"`           // 消息内容
	MsgType      string    `gorm:"column:msg_type" json:"msg_type"`         // 消息类型
	MsgSubType   string    `gorm:"column:msg_sub_type" json:"msg_sub_type"` // 消息子类型
	MsgTime      time.Time `gorm:"column:msg_time" json:"msg_time"`         // 消息时间
	IsDeleted    bool      `gorm:"column:is_deleted" json:"is_deleted"`     // 是否删除
	IsSend       bool      `gorm:"column:is_send" json:"is_send"`           // 是否是工作微信发出的消息
	SendStatus   bool      `gorm:"column:send_status" json:"send_status"`   // 消息发送状态
	Origin       string    `gorm:"column:origin" json:"origin"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *WxMsg) TableName() string {
	return "wx_msg"
}
