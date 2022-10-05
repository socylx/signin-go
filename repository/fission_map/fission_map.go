package fission_map

import (
	"time"
)

// FissionMap 拉新关系表
type FissionMap struct {
	ID            uint32    `gorm:"primaryKey;column:id" json:"id"`
	ShareUserID   uint32    `gorm:"column:share_user_id" json:"share_user_id"`     // 分享人id, 发送邀请的人
	InvitedUserID uint32    `gorm:"column:invited_user_id" json:"invited_user_id"` // 被邀请人id
	Type          uint32    `gorm:"column:type" json:"type"`                       // 类型，1老学员邀请新学员，2工作人员邀请老学员
	Status        uint32    `gorm:"column:status" json:"status"`                   // 状态，1邀请还未接受，2邀请接受，3已购买体验券，4已上课，5已经办卡，6老用户（无效邀请），7已经被别人邀请（无效邀请）8邀请失败，已经有优惠券了
	IsDel         bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	ActivityMark  uint32    `gorm:"column:activity_mark" json:"activity_mark"` // 小猫咪活动标记，0没有活动，1双十一
	Mark          uint32    `gorm:"column:mark" json:"mark"`                   // 1-正常, 2-申请回复奖励
}

type FissionMapData struct {
	ID uint32 ` bson:"id" json:"id"`
}
