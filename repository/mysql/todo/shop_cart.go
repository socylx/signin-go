package todo

import (
	"time"
)

// ShopCart 购物车
type ShopCart struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID     uint32    `gorm:"column:user_id" json:"user_id"` // 用户id
	Type       uint32    `gorm:"column:type" json:"type"`       // 商品类别，1商品，2课程
	SkuID      uint32    `gorm:"column:sku_id" json:"sku_id"`
	Num        uint32    `gorm:"column:num" json:"num"` // 商品数量
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopCart) TableName() string {
	return "shop_cart"
}
