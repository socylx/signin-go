package user_snapshot

import (
	"time"
)

// UserSnapshot [...]
type UserSnapshot struct {
	ID                uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID            uint32    `gorm:"column:user_id" json:"user_id"`                         // 用户ID
	UserType          uint32    `gorm:"column:user_type" json:"user_type"`                     // 客户类型，成人/少儿
	MembershipRemains float32   `gorm:"column:membership_remains" json:"membership_remains"`   // 剩余会员卡次数
	CouponRemains     float32   `gorm:"column:coupon_remains" json:"coupon_remains"`           // 剩余券次数
	IntervalStage     uint32    `gorm:"column:interval_stage" json:"interval_stage"`           // 区间阶段，1:p1阶段，2:p2阶段， 3:p3阶段...
	MembershipIDs     string    `gorm:"column:membership_ids" json:"membership_ids"`           // 多个可用会员卡ID字符串，形如：1:1,2:10, 代表ID：1的会员卡还有1次，ID：2的会员卡还有10次
	CouponAllocIDs    string    `gorm:"column:coupon_alloc_ids" json:"coupon_alloc_ids"`       // 多个可用券ID字符串，形如：4:2,8:1
	CreateTime        time.Time `gorm:"column:create_time" json:"create_time"`                 // 创建时间
	UpdateTime        time.Time `gorm:"column:update_time" json:"update_time"`                 // 更新时间
	IsDel             bool      `gorm:"column:is_del" json:"is_del"`                           // 是否删除
	RenewMembershipID uint32    `gorm:"column:renew_membership_id" json:"renew_membership_id"` // 如果在这个节点续卡了，续的卡id
}

// TableName get sql table name.获取数据库表名
func tableName() string {
	return "user_snapshot"
}
