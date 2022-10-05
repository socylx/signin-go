package users

import "signin-go/global/time"

type StrategyIndicatorScore struct {
	DocumentID         interface{}            `bson:"_id,omitempty" json:"_id"`
	UserID             uint32                 `bson:"user_id" json:"user_id"`
	UserBeforeMemberID uint32                 `bson:"user_before_member_id" json:"user_before_member_id"`
	Time               time.Time              `bson:"time" json:"time"`
	Type               uint32                 `bson:"type" json:"type"`
	StrategyKey        string                 `bson:"strategy_key" json:"strategy_key"`
	Data               map[string]interface{} `bson:"data" json:"data"`
	Scores             []*Score               `bson:"scores" json:"scores"`
}

type Score struct {
	ID     uint32  `bson:"id" json:"id"`
	Name   string  `bson:"name" json:"name"`
	Score  float64 `bson:"score" json:"score"`
	Weight uint32  `bson:"weight" json:"weight"`
}
