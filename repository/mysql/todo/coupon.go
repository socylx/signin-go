package todo

import (
	"time"
)

// Coupon 券模板
type Coupon struct {
	ID             uint32    `gorm:"primaryKey;column:id" json:"id"`
	Type           int16     `gorm:"column:type" json:"type"`               // 类型：1-满减券；2-代金券；3-折扣券
	AmountType     int16     `gorm:"column:amount_type" json:"amount_type"` // 价值类型：1-金额，2次数，3-时长
	Amount         float32   `gorm:"column:amount" json:"amount"`           // 额度（元/次/小时）
	Name           string    `gorm:"column:name" json:"name"`               // 优惠券名称
	OneOff         bool      `gorm:"column:one_off" json:"one_off"`         // 一次性使用即失效
	Description    string    `gorm:"column:description" json:"description"` // 描述信息
	Mark           string    `gorm:"column:mark" json:"mark"`               // 备注信息
	Expire         uint32    `gorm:"column:expire" json:"expire"`           // 有效时间（小时），默认三年
	AutoEffect     bool      `gorm:"column:auto_effect" json:"auto_effect"` // 是否自动生效，不需要领券，只要满足条件就自动执行。0不是，1是
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel          uint8     `gorm:"column:is_del" json:"is_del"`                     // 逻辑删除标识，0-有效数据，1-无效数据
	NeedPay        bool      `gorm:"column:need_pay" json:"need_pay"`                 // 是否需要购买
	Price          float32   `gorm:"column:price" json:"price"`                       // 售价
	PerUserSum     uint32    `gorm:"column:per_user_sum" json:"per_user_sum"`         // 每个用户拥有这个券的最大数量
	IsNewUser      bool      `gorm:"column:is_new_user" json:"is_new_user"`           // 是否是新人券（标识新人）0不是，1是
	CanUserOpt     bool      `gorm:"column:can_user_opt" json:"can_user_opt"`         // 该券用户是否能自己领取。0只能系统或者管理管理员发放
	GlobalLimitNum uint32    `gorm:"column:global_limit_num" json:"global_limit_num"` // 该券整个系统限制发放数量，0不限制
	UserType       uint32    `gorm:"column:user_type" json:"user_type"`               // 限制获取的用户类型，0不限制，1有效会员，2失效会员，3所有会员，4新人，5非会员
}

// TableName get sql table name.获取数据库表名
func (m *Coupon) TableName() string {
	return "coupon"
}
