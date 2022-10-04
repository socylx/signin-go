package strategy_indicator_rule

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

func GetStrategyIndicatorRules(ctx core.StdContext, IDs []uint32) (data []*StrategyIndicatorRule, err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table("strategy_indicator_rule").
		Select("strategy_indicator_rule.*").
		Where("strategy_indicator_rule.is_del = 0 AND strategy_indicator_rule.id IN ?", IDs).
		Find(&data).Error
	return
}
