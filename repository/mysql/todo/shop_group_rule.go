package todo

import (
	"time"
)

// ShopGroupRule [...]
type ShopGroupRule struct {
	ID          uint32    `gorm:"primaryKey;column:id" json:"id"`
	ShopGroupID uint32    `gorm:"column:shop_group_id" json:"shop_group_id"`
	Type        uint32    `gorm:"column:type" json:"type"` // 1发起人类型，2参团人类型
	Value       uint32    `gorm:"column:value" json:"value"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel       bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopGroupRule) TableName() string {
	return "shop_group_rule"
}
