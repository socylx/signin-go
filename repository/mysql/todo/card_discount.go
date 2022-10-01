package todo

import (
	"time"
)

// CardDiscount [...]
type CardDiscount struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	ForID      uint32    `gorm:"column:for_id" json:"for_id"`
	ForType    uint32    `gorm:"column:for_type" json:"for_type"` // 1-系列课里的课程
	CardID     uint32    `gorm:"column:card_id" json:"card_id"`   // 卡类型ID
	Amount     float32   `gorm:"column:amount" json:"amount"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *CardDiscount) TableName() string {
	return "card_discount"
}
