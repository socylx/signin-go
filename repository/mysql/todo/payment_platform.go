package todo

import (
	"time"
)

// PaymentPlatform 收款平台
type PaymentPlatform struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"` // 名称，自有微信支付、招行平台、自有支付宝平台、现金
	Key        string    `gorm:"column:key" json:"key"`   // 平台标记，self_wechat, self_alipay,cmb,cash
	Mark       string    `gorm:"column:mark" json:"mark"` // 备注
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *PaymentPlatform) TableName() string {
	return "payment_platform"
}
