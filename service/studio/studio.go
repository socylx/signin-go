package studio

import (
	"signin-go/internal/core"
	"signin-go/repository/studio_strategy_map"
)

type StudioStrategyTypeIDMap map[uint32]map[uint32]uint32

/*
获取门店应用的某种类型的策略对应的策略ID
*/
func GetStudioStrategyTypeIDMap(ctx core.StdContext, studioIDs []uint32) (data StudioStrategyTypeIDMap, err error) {
	studioStrategyMapDatas, err := studio_strategy_map.GetStudioStrategyMapDatas(ctx, studioIDs)
	if err != nil {
		return
	}
	data = StudioStrategyTypeIDMap{}
	for _, studioStrategyMapData := range studioStrategyMapDatas {
		strategyTypeIDMap, ok := data[studioStrategyMapData.StudioID]
		if !ok {
			strategyTypeIDMap = map[uint32]uint32{}
		}
		strategyTypeIDMap[studioStrategyMapData.StrategyType] = studioStrategyMapData.StrategyID
		data[studioStrategyMapData.StudioID] = strategyTypeIDMap
	}
	return
}
