package strategy_indicator

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

func List(ctx core.StdContext) (data []*StrategyIndicator, err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table("strategy_indicator").
		Where("strategy_indicator.is_del = 0").
		Find(&data).Error
	return
}
