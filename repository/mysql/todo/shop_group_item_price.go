package todo

import (
	"time"
)

// ShopGroupItemPrice [...]
type ShopGroupItemPrice struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"id"`
	ShopGroupItemID uint32    `gorm:"column:shop_group_item_id" json:"shop_group_item_id"`
	Price           float32   `gorm:"column:price" json:"price"`
	MasterPrice     float32   `gorm:"column:master_price" json:"master_price"` // 团长价格
	Type            uint32    `gorm:"column:type" json:"type"`                 // 1金额，2卡次
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopGroupItemPrice) TableName() string {
	return "shop_group_item_price"
}
