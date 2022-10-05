package strategy

import (
	"signin-go/repository/strategy_indicator"
	"signin-go/repository/strategy_indicator_rule"
	"time"
)

// Strategy [...]
type Strategy struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name         string    `gorm:"column:name" json:"name"`
	Desc         string    `gorm:"column:desc" json:"desc"`
	Type         uint32    `gorm:"column:type" json:"type"` // 1-拉新，2-续卡
	Status       bool      `gorm:"column:status" json:"status"`
	StartTime    time.Time `gorm:"column:start_time" json:"start_time"` // 开始时间
	Key          string    `gorm:"column:key" json:"key"`
	CreateUserID uint32    `gorm:"column:create_user_id" json:"create_user_id"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
}

type IndicatorData struct {
	strategy_indicator.StrategyIndicator
	Weight uint32               `json:"weight"`
	Rules  []*IndicatorRuleData `json:"rules"`
}

type IndicatorRuleData struct {
	strategy_indicator_rule.StrategyIndicatorRule
	Score uint32 `json:"score"`
}