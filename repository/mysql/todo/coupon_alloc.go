package todo

import (
	"time"
)

// CouponAlloc 券实例
type CouponAlloc struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	OrderID    uint32    `gorm:"column:order_id" json:"order_id"`       // 如果通过订单购买，为订单ID
	CouponID   uint32    `gorm:"column:coupon_id" json:"coupon_id"`     // 优惠券ID
	Remain     float32   `gorm:"column:remain" json:"remain"`           // 该券的剩余余额
	UserID     uint32    `gorm:"column:user_id" json:"user_id"`         // 用户ID
	CouponCode string    `gorm:"column:coupon_code" json:"coupon_code"` // 券号
	Status     bool      `gorm:"column:status" json:"status"`           // 使用状态：0-未使用；1-已使用；2失效，3未激活
	Deadline   time.Time `gorm:"column:deadline" json:"deadline"`       // 截止时间
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"` // 操作人id, 例如修改了这张券的状态的人
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 发放时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`           // 是否有效：0-有效；1-无效
	GetType    uint32    `gorm:"column:get_type" json:"get_type"`       // 获取方式：1订单购买，2管理员发放，3免费领取，4老学员拉新赠送，5老带新每月折扣活动
	OtherCode  string    `gorm:"column:other_code" json:"other_code"`   // 其他平台的码，例如大众点评购买的体验券码
	Mark       string    `gorm:"column:mark" json:"mark"`               // 发券备注
}

// TableName get sql table name.获取数据库表名
func (m *CouponAlloc) TableName() string {
	return "coupon_alloc"
}
