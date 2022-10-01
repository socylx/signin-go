package todo

import (
	"time"
)

// ShopGroupItem [...]
type ShopGroupItem struct {
	ID                 uint32    `gorm:"primaryKey;column:id" json:"id"`
	ShopGroupID        uint32    `gorm:"column:shop_group_id" json:"shop_group_id"` // 所属拼团活动id
	Type               uint32    `gorm:"column:type" json:"type"`                   // 1会员卡，2课程，3优惠券, 4sku 商品1会员卡，2课程，3优惠券, 4sku 商品
	ForID              uint32    `gorm:"column:for_id" json:"for_id"`               // 对应的商品id
	MinNum             uint32    `gorm:"column:min_num" json:"min_num"`             // 最小购买数量
	MaxNum             uint32    `gorm:"column:max_num" json:"max_num"`             // 最大购买数量
	IsDel              bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime         time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime         time.Time `gorm:"column:create_time" json:"create_time"`
	AutoGroup          bool      `gorm:"column:auto_group" json:"auto_group"`                     // 是否智能开团
	AutoGroupTimedelta uint32    `gorm:"column:auto_group_timedelta" json:"auto_group_timedelta"` // 智能开团的间隔
	AutoJoin           bool      `gorm:"column:auto_join" json:"auto_join"`                       // 是否智能参团
}

// TableName get sql table name.获取数据库表名
func (m *ShopGroupItem) TableName() string {
	return "shop_group_item"
}
