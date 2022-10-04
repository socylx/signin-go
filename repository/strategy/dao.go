package strategy

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

type DeleteFileter struct {
	ID uint32
}

func Delete(ctx core.StdContext, filter *DeleteFileter) (err error) {
	if filter.ID <= 0 {
		return
	}

	db := mysql.DB.WithContext(ctx)
	err = db.Table(tableName()).
		Where("strategy.is_del = 0 AND strategy.id = ?", filter.ID).
		Update("strategy.is_del", 1).Error

	return
}

func Detail(ctx core.StdContext, strategyID uint32) (strategy *Strategy, err error) {
	db := mysql.DB.WithContext(ctx)

	strategy = &Strategy{}
	err = db.Table(tableName()).
		Where("strategy.is_del = 0 AND strategy.id = ?", strategyID).
		First(strategy).Error
	return
}

func GetStrategyIndicatorDatas(ctx core.StdContext, strategyIndicatorIDs []uint32) (data []*IndicatorData, err error) {
	db := mysql.DB.WithContext(ctx)
	db.Table("strategy_indicator").
		Select("strategy_indicator.*").
		Where("strategy_indicator.is_del = 0 AND strategy_indicator.id IN ?", strategyIndicatorIDs).
		Find(&data)
	return
}

func GetStrategyIndicatorRuleDatas(ctx core.StdContext, strategyIndicatorRuleIDs []uint32) (data []*IndicatorRuleData, err error) {
	db := mysql.DB.WithContext(ctx)

	err = db.Table("strategy_indicator_rule").
		Select("strategy_indicator_rule.*").
		Where("strategy_indicator_rule.is_del = 0 AND strategy_indicator_rule.id IN ?", strategyIndicatorRuleIDs).
		Find(&data).Error
	return
}
