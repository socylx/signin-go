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
