package todo

import (
	"time"
)

// MembershipShare [...]
type MembershipShare struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	MembershipID uint32    `gorm:"column:membership_id" json:"membership_id"`
	StudioID     uint32    `gorm:"column:studio_id" json:"studio_id"`
	UserID       uint32    `gorm:"column:user_id" json:"user_id"`
	Amount       float32   `gorm:"column:amount" json:"amount"`
	Remark       string    `gorm:"column:remark" json:"remark"`
	OptUserID    uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *MembershipShare) TableName() string {
	return "membership_share"
}
