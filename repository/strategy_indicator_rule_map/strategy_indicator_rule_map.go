package strategy_indicator_rule_map

import (
	"time"
)

// StrategyIndicatorRuleMap [...]
type StrategyIndicatorRuleMap struct {
	ID                      uint32    `gorm:"primaryKey;column:id" json:"id"`
	StrategyID              uint32    `gorm:"column:strategy_id" json:"strategy_id"`
	StrategyIndicatorRuleID uint32    `gorm:"column:strategy_indicator_rule_id" json:"strategy_indicator_rule_id"`
	Weight                  uint32    `gorm:"column:weight" json:"weight"`
	Score                   uint32    `gorm:"column:score" json:"score"`
	CreateTime              time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime              time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel                   bool      `gorm:"column:is_del" json:"is_del"`
}
