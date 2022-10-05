package strategy_indicator_rule_map

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

func GetStrategyIndicatorRuleMaps(ctx core.StdContext, strategyID uint32) (data []*StrategyIndicatorRuleMap, err error) {
	db := mysql.DB.WithContext(ctx)

	err = db.Table("strategy_indicator_rule_map").
		Where("strategy_indicator_rule_map.is_del = 0 AND strategy_indicator_rule_map.strategy_id = ?", strategyID).
		Find(&data).Error
	return
}

func Delete(ctx core.StdContext, strategyID uint32) (err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table("strategy_indicator_rule_map").
		Where("strategy_indicator_rule_map.is_del = 0 AND strategy_indicator_rule_map.strategy_id = ?", strategyID).
		Update("strategy_indicator_rule_map.is_del", 1).Error
	return
}
