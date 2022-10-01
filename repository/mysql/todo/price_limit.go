package todo

import (
	"time"
)

// PriceLimit 会员卡的使用限制，例如限制会员卡套餐，只能在指定的门店使用
type PriceLimit struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	PriceID    uint32    `gorm:"column:price_id" json:"price_id"`       // 对应的限制会员卡 price id
	Value      string    `gorm:"column:value" json:"value"`             // 会员卡的限制值,例如现在会员卡只能用于具体门店，这个值就代表门店id
	Type       int       `gorm:"column:type" json:"type"`               // 限制类型，1限制门店
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`           // 是否已被删除
}

// TableName get sql table name.获取数据库表名
func (m *PriceLimit) TableName() string {
	return "price_limit"
}
