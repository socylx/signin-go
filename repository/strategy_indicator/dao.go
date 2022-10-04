package strategy_indicator

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
	"signin-go/repository/strategy_indicator_rule"
)

type StrategyIndicatorData struct {
	StrategyIndicator
	StrategyIndicatorRules []*strategy_indicator_rule.StrategyIndicatorRule `gorm:"ForeignKey:StrategyIndicatorID;AssociationForeignKey:StrategyIndicatorID" json:"rules"`
}

func List(ctx core.StdContext) (data []*StrategyIndicatorData, err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table(tableName()).
		Preload("StrategyIndicatorRules", "strategy_indicator_rule.is_del = 0").
		Where("strategy_indicator.is_del = 0").
		Find(&data).Error
	return
}
