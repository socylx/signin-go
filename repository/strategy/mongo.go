package strategy

import (
	"time"
)

const StrategyName = "strategy"

type StrategyDocument struct {
	DocumentID   interface{} `bson:"_id,omitempty" json:"_id"`
	ID           uint32      `bson:"id"`
	Name         string      `bson:"name"`
	Desc         string      `bson:"desc"`
	Type         uint32      `bson:"type"` // 1-拉新，2-续卡
	Status       bool        `bson:"status"`
	StartTime    time.Time   `bson:"start_time"` // 开始时间
	Key          string      `bson:"key"`
	CreateUserID uint32      `bson:"create_user_id"`
	CreateTime   time.Time   `bson:"create_time"`
	UpdateTime   time.Time   `bson:"update_time"`
	IsDel        bool        `bson:"is_del"`

	StrategyIndicators []*StrategyIndicator `bson:"strategy_indicators"`
}

type StrategyIndicator struct {
	ID                     uint32                   `bson:"id"`
	Key                    string                   `bson:"key"`
	Name                   string                   `bson:"name"`
	CreateTime             time.Time                `bson:"create_time"`
	UpdateTime             time.Time                `bson:"update_time"`
	IsDel                  bool                     `bson:"is_del"`
	Weight                 uint32                   `bson:"weight"`
	StrategyIndicatorRules []*StrategyIndicatorRule `bson:"strategy_indicator_rules"`
}

type StrategyIndicatorRule struct {
	ID         uint32    `bson:"id"`
	Name       string    `bson:"name"`
	Min        string    `bson:"min"`
	Max        string    `bson:"max"`
	CreateTime time.Time `bson:"create_time"`
	UpdateTime time.Time `bson:"update_time"`
	IsDel      bool      `bson:"is_del"`
	Score      uint32    `bson:"score"`
}
