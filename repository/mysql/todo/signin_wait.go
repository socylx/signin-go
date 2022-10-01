package todo

import (
	"time"
)

// SigninWait 约课排队表
type SigninWait struct {
	ID            uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID        uint32    `gorm:"column:user_id" json:"user_id"`                 // 对应系统注册用户id
	ActivityID    uint32    `gorm:"column:activity_id" json:"activity_id"`         // 等待的课程id
	CourseID      uint32    `gorm:"column:course_id" json:"course_id"`             // 课程组id
	IsBundling    bool      `gorm:"column:is_bundling" json:"is_bundling"`         // 课程是否是绑定预约的
	Status        uint32    `gorm:"column:status" json:"status"`                   // 等待状态：1等待中，2取消等待，3约课成功，4约课失败
	MembershipID  uint32    `gorm:"column:membership_id" json:"membership_id"`     // 支付的会员卡id
	CouponAllocID uint32    `gorm:"column:coupon_alloc_id" json:"coupon_alloc_id"` // 用来支付的券id
	Mark          string    `gorm:"column:mark" json:"mark"`                       // 备注, 失败原因
	IsNotice      uint32    `gorm:"column:is_notice" json:"is_notice"`             // 发送结果通知，0未发送，1已经发送
	OptUserID     uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`         // 操作人id
	IsDel         bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
	SigninTime    time.Time `gorm:"column:signin_time" json:"signin_time"`     // 约课成功时间
	IndexTime     time.Time `gorm:"column:index_time" json:"index_time"`       // 排课时间
	Type          uint32    `gorm:"column:type" json:"type"`                   // 1-正常排队，2-取消约课转入排队
	DelayMinutes  uint32    `gorm:"column:delay_minutes" json:"delay_minutes"` // 提前多少分钟，可以让位给我
}

// TableName get sql table name.获取数据库表名
func (m *SigninWait) TableName() string {
	return "signin_wait"
}
