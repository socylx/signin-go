package todo

import (
	"time"
)

// OrderPayment 订单的支付单，一个订单可以分为多次支付
type OrderPayment struct {
	ID               uint32    `gorm:"primaryKey;column:id" json:"id"`
	OrderID          uint32    `gorm:"column:order_id" json:"order_id"`                     // 对应的订单
	Status           uint16    `gorm:"column:status" json:"status"`                         // 状态，1待支付，2已支付，3失效，4废弃
	OuterPayNo       string    `gorm:"column:outer_pay_no" json:"outer_pay_no"`             // 外部的账单号
	TradeNo          string    `gorm:"column:trade_no" json:"trade_no"`                     // 商户系统内部付款订单号
	Total            float64   `gorm:"column:total" json:"total"`                           // 合计价格
	Discount         float64   `gorm:"column:discount" json:"discount"`                     // 折扣
	PaymentAccountID uint32    `gorm:"column:payment_account_id" json:"payment_account_id"` // 收款账号
	PayType          uint32    `gorm:"column:pay_type" json:"pay_type"`                     // 实际付款方式，1微信支付，2支付宝，3银联，4现金，5支付宝分期
	PayBank          uint32    `gorm:"column:pay_bank" json:"pay_bank"`                     // 付款银行
	PayTime          time.Time `gorm:"column:pay_time" json:"pay_time"`                     // 支付完成时间
	Mark             string    `gorm:"column:mark" json:"mark"`                             // 备注
	SysMark          string    `gorm:"column:sys_mark" json:"sys_mark"`                     // 系统备注
	IsDel            bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *OrderPayment) TableName() string {
	return "order_payment"
}
