package strategy

import (
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
	ID         uint32               `bson:"id" json:"id"`
	Key        string               `bson:"key" json:"key"`
	Name       string               `bson:"name" json:"name"`
	CreateTime time.Time            `bson:"create_time" json:"create_time"`
	UpdateTime time.Time            `bson:"update_time" json:"update_time"`
	IsDel      bool                 `bson:"is_del" json:"is_del"`
	Weight     uint32               `bson:"weight" json:"weight"`
	Rules      []*IndicatorRuleData `bson:"rules" json:"rules"`
}

type IndicatorRuleData struct {
	strategy_indicator_rule.StrategyIndicatorRule
	Score uint32 `json:"score"`
}
