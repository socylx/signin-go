package studio_strategy_map

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

func GetStudioStrategyIDs(ctx core.StdContext, studioIDs []uint32) (data []uint32, err error) {
	db := mysql.DB.WithContext(ctx)

	err = db.Table("studio_strategy_map").
		Select("studio_strategy_map.strategy_id").
		Where("studio_strategy_map.is_del = 0 AND studio_strategy_map.studio_id IN ?", studioIDs).
		Find(&data).Error
	return
}

func GetStudioStrategyMapDatas(ctx core.StdContext, studioIDs []uint32) (data []*StudioStrategyMapData, err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table("studio_strategy_map").
		Select("studio_strategy_map.*, strategy.type as strategy_type, strategy.status as strategy_status").
		Joins("JOIN strategy on studio_strategy_map.strategy_id = strategy.id").
		Where("studio_strategy_map.is_del = 0 AND studio_strategy_map.studio_id IN ?", studioIDs).
		Where("strategy.is_del = 0 AND strategy.status = 1").
		Find(&data).Error
	return
}
