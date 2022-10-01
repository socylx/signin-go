package todo

import (
	"time"
)

// DuizhangOrder [...]
type DuizhangOrder struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	DuizhangID uint32    `gorm:"column:duizhang_id" json:"duizhang_id"`
	OrderID    uint32    `gorm:"column:order_id" json:"order_id"`
	ForID      uint32    `gorm:"column:for_id" json:"for_id"`
	ForType    uint32    `gorm:"column:for_type" json:"for_type"` // 0: 无, 1: order_payment, 2: order_refund
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *DuizhangOrder) TableName() string {
	return "duizhang_order"
}
