package todo

import (
	"time"
)

// WxFriend [...]
type WxFriend struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	StaffUserID  uint32    `gorm:"column:staff_user_id" json:"staff_user_id"`   // 工作人员user_id
	FriendUserID uint32    `gorm:"column:friend_user_id" json:"friend_user_id"` // 学员user_id
	UserBeforeID uint32    `gorm:"column:user_before_id" json:"user_before_id"`
	ConRemark    string    `gorm:"column:con_remark" json:"con_remark"`
	IsPassed     bool      `gorm:"column:is_passed" json:"is_passed"`   // 是否通过好友请求
	IsDeleted    bool      `gorm:"column:is_deleted" json:"is_deleted"` // 是否删除了此好友
	AddTime      time.Time `gorm:"column:add_time" json:"add_time"`     // 请求添加时间
	AddFrom      string    `gorm:"column:add_from" json:"add_from"`     // 添加来源
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *WxFriend) TableName() string {
	return "wx_friend"
}
