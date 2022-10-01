package todo

import (
	"time"
)

// ShopPromotionItem 促销活动的商品（课程或者卡等）
type ShopPromotionItem struct {
	ID          uint32    `gorm:"primaryKey;column:id" json:"id"`
	PromotionID uint32    `gorm:"column:promotion_id" json:"promotion_id"` // 所属促销id
	ItemID      uint32    `gorm:"column:item_id" json:"item_id"`           // 商品id
	Type        uint32    `gorm:"column:type" json:"type"`                 // 商品类型，1课程，2商品
	IsDel       bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopPromotionItem) TableName() string {
	return "shop_promotion_item"
}
