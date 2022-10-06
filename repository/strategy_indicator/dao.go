package strategy_indicator

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

func List(ctx core.StdContext, IDs []uint32) (data []*StrategyIndicator, err error) {
	db := mysql.DB.WithContext(ctx)
	query := db.Table("strategy_indicator").Where("strategy_indicator.is_del = 0")
	if len(IDs) > 0 {
		query = query.Where("strategy_indicator.id IN ?", IDs)
	}
	err = query.Find(&data).Error
	return
}
