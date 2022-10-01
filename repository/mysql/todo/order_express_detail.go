package todo

import (
	"time"
)

// OrderExpressDetail [...]
type OrderExpressDetail struct {
	ID             uint32    `gorm:"primaryKey;column:id" json:"id"`
	OrderExpressID uint32    `gorm:"column:order_express_id" json:"order_express_id"` // 快递ID
	OrderItemID    uint32    `gorm:"column:order_item_id" json:"order_item_id"`       // 快递包含的商品
	Num            uint32    `gorm:"column:num" json:"num"`                           // 快递发的商品数量
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel          bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *OrderExpressDetail) TableName() string {
	return "order_express_detail"
}
