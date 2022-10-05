package order_refund_item

import (
	"time"
)

// OrderRefundItem [...]
type OrderRefundItem struct {
	ID                 uint32    `gorm:"primaryKey;column:id" json:"id"`
	OrderRefundID      uint32    `gorm:"column:order_refund_id" json:"order_refund_id"`             // 退款的商品所属的退款记录ID
	OrderItemPrimaryID uint32    `gorm:"column:order_item_primary_id" json:"order_item_primary_id"` // order_item 表的id
	OrderItemID        uint32    `gorm:"column:order_item_id" json:"order_item_id"`                 // 退款的商品ID, sku_id。
	InstanceID         uint32    `gorm:"column:instance_id" json:"instance_id"`                     // 实例Id
	ItemNum            uint32    `gorm:"column:item_num" json:"item_num"`                           // 退款商品数量, 为0时部分退款
	Amount             float32   `gorm:"column:amount" json:"amount"`                               // 退款金额, 单位：元
	Type               uint32    `gorm:"column:type" json:"type"`                                   // 退款类型, 1-订单退款, 2-退手续费
	CreateTime         time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime         time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel              bool      `gorm:"column:is_del" json:"is_del"`
}
