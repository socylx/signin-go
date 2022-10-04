package studio_strategy_map

import (
	"time"
)

// StudioStrategyMap [...]
type StudioStrategyMap struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	StudioID   uint32    `gorm:"column:studio_id" json:"studio_id"`
	StrategyID uint32    `gorm:"column:strategy_id" json:"strategy_id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}
