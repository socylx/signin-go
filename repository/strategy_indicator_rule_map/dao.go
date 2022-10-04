package strategy_indicator_rule_map

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

func GetStrategyIndicatorRuleMaps(ctx core.StdContext, strategyID uint32) (data []*StrategyIndicatorRuleMap, err error) {
	db := mysql.DB.WithContext(ctx)

	err = db.Table(tableName()).
		Where("strategy_indicator_rule_map.is_del = 0 AND strategy_indicator_rule_map.strategy_id = ?", strategyID).
		Find(&data).Error
	return
}
