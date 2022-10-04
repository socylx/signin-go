package strategy_indicator

import (
	"time"
)

// StrategyIndicator [...]
type StrategyIndicator struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Key        string    `gorm:"column:key" json:"key"`
	Name       string    `gorm:"column:name" json:"name"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}
