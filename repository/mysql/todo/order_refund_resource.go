package todo

import (
	"time"
)

// OrderRefundResource [...]
type OrderRefundResource struct {
	ID            uint32    `gorm:"primaryKey;column:id" json:"id"`
	OrderRefundID uint32    `gorm:"column:order_refund_id" json:"order_refund_id"`
	URL           string    `gorm:"column:url" json:"url"`
	Type          uint32    `gorm:"column:type" json:"type"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel         bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *OrderRefundResource) TableName() string {
	return "order_refund_resource"
}
