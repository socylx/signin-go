package todo

import (
	"time"
)

// MembershipTransform [...]
type MembershipTransform struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"id"`
	OldMembershipID uint32    `gorm:"column:old_membership_id" json:"old_membership_id"` // 更换的老卡id
	OldAmount       float32   `gorm:"column:old_amount" json:"old_amount"`               // 老卡剩余金额
	OldRemains      float32   `gorm:"column:old_remains" json:"old_remains"`             // 老卡的剩余次数
	NewMembershipID uint32    `gorm:"column:new_membership_id" json:"new_membership_id"` // 更换的新卡id
	NewAmount       float32   `gorm:"column:new_amount" json:"new_amount"`               // 新卡补款金额
	NewRemains      float32   `gorm:"column:new_remains" json:"new_remains"`             // 新卡的剩余次数
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`
	Mark            string    `gorm:"column:mark" json:"mark"`               // 备注
	OptUserID       uint32    `gorm:"column:opt_user_id" json:"opt_user_id"` // 操作人id
}

// TableName get sql table name.获取数据库表名
func (m *MembershipTransform) TableName() string {
	return "membership_transform"
}
