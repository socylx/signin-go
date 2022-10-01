package todo

import (
	"time"
)

// OrderItem2 [...]
type OrderItem2 struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	OrderID    uint32    `gorm:"column:order_id" json:"order_id"`
	ItemID     uint32    `gorm:"column:item_id" json:"item_id"` // 购买的商品id, price_id
	Type       uint32    `gorm:"column:type" json:"type"`       // 商品的类型，1会员卡，2课程
	Status     uint16    `gorm:"column:status" json:"status"`   // 订单下商品的状态，0未处理，4已完成
	Price      float32   `gorm:"column:price" json:"price"`     // 商品单价
	Num        int       `gorm:"column:num" json:"num"`         // 购买数量
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
	Mark       string    `gorm:"column:mark" json:"mark"`               // 备注，记录，备用字段
	Number     string    `gorm:"column:number" json:"number"`           // 编号，若是会员卡，对应的是实体卡卡号
	InstanceID uint32    `gorm:"column:instance_id" json:"instance_id"` // 实例Id
}

// TableName get sql table name.获取数据库表名
func (m *OrderItem2) TableName() string {
	return "order_item2"
}
