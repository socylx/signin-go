package order_refund

import (
	"time"
)

// OrderRefund 订单退款
type OrderRefund struct {
	ID               uint32    `gorm:"primaryKey;column:id" json:"id"`
	OrderID          uint32    `gorm:"column:order_id" json:"order_id"`                     // 退款对应的订单
	PaymentAccountID uint32    `gorm:"column:payment_account_id" json:"payment_account_id"` // 如果是走的公用的退款账户：支付这次退款的，付款账户id
	OrderPaymentID   uint32    `gorm:"column:order_payment_id" json:"order_payment_id"`     // 如果走的是原路退回：支付id, 如果是按原路退款的，则有这个id
	OutRefundNo      string    `gorm:"column:out_refund_no" json:"out_refund_no"`           // 退款订单号
	RefundID         string    `gorm:"column:refund_id" json:"refund_id"`                   // 微信官方交易退款单号
	Status           uint32    `gorm:"column:status" json:"status"`                         // 1: 发起退款, 2：退款成功, 3: 退款关闭, 4: 退款异常
	SuccessTime      time.Time `gorm:"column:success_time" json:"success_time"`             // 成功退款时间
	RecvAccout       string    `gorm:"column:recv_accout" json:"recv_accout"`               // 退款入账方
	Remark           string    `gorm:"column:remark" json:"remark"`                         // 退款备注
	SysRemark        string    `gorm:"column:sys_remark" json:"sys_remark"`                 // 工作人员退款备注
	OptUserID        uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`               // 退款操作人
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel            bool      `gorm:"column:is_del" json:"is_del"`
}
