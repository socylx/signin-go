package todo

import (
	"time"
)

// CouponRule 使用券的规则限制，例如限制某些课程组课使用
type CouponRule struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	CouponID   uint32    `gorm:"column:coupon_id" json:"coupon_id"` // 券模板id
	Type       uint32    `gorm:"column:type" json:"type"`           // 限制的类型，1限制课程类型（大师课，小班课），2限制课程组，3限制具体课程，4金额（大于等于），5金额（小于等于），6限制门店，7限制会员类型，8有效日期（大于等于），9有效日期（小于等于）,10 限制课程消费的卡类型（例如只有用于少儿卡消费的课程），11 促销活动聚合，12 SPU商品，13 SKU商品，14上课方式，15商品数量（大于等于），16商品数量（小于等于）
	Value      string    `gorm:"column:value" json:"value"`         // 对应限制的值，可能是课程id, 金额等。类型为字符串，逻辑处理时需要转换；如果是限制会员类型：1老会员，2新会员
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *CouponRule) TableName() string {
	return "coupon_rule"
}
