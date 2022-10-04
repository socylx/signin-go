package strategy

import (
	"signin-go/internal/code"
	"signin-go/internal/core"
	strategyRepo "signin-go/repository/strategy"
	studioRepo "signin-go/repository/studio"
	"signin-go/repository/user_before_member"
	studioServ "signin-go/service/studio"
	"sync"
)

func GenerateOfLaxin(ctx core.StdContext) core.BusinessError {
	studioIDs, err := studioRepo.GetStudioIDs(ctx)
	if err != nil || len(studioIDs) <= 0 {
		return core.Error(code.StudioQueryError, code.Text(code.StudioQueryError)).WithError(err)
	}

	studioStrategyTypeIDMap, err := studioServ.GetStudioStrategyTypeIDMap(ctx, studioIDs)
	if err != nil {
		return core.Error(code.StudioQueryError, "获取门店应用的某种类型的策略对应的策略ID").WithError(err)
	}

	strategyIndicatorsDataMap := map[uint32]*strategyRepo.StrategyDocument{}
	for _, strategyTypeID := range studioStrategyTypeIDMap {
		for _, strategyID := range strategyTypeID {
			strategyIndicatorsData, err := Data(ctx, strategyID)
			if err != nil {
				continue
			}
			strategyIndicatorsDataMap[strategyID] = strategyIndicatorsData
		}
	}

	maxUserBeforeMemberID, err := user_before_member.GetMaxUserBeforeMemberID(ctx)
	if err != nil {
		return core.Error(code.UserBeforeMemberQueryError, "查询线索最大值失败").WithError(err)
	}

	var wg sync.WaitGroup
	var lastUserBeforeMemberID int64
	for {
		if lastUserBeforeMemberID > maxUserBeforeMemberID {
			break
		}
		userBeforeMemberIDs, err := user_before_member.GetUserBeforeMemberIDs(ctx, &user_before_member.Filter{IDGT: lastUserBeforeMemberID})
		lastUserBeforeMemberID += 1000
		if err != nil || len(userBeforeMemberIDs) <= 0 {
			continue
		}

		// wg.Add(1)
		// go func(userBeforeMemberIDs []int64) {
		// 	for _, userBeforeMemberID := range userBeforeMemberIDs {
		// 		userData, err := h.userService.Data(c.RequestContext(), &user.DataFilter{
		// 			UserID:             0,
		// 			UserBeforeMemberID: uint32(userBeforeMemberID),
		// 		})
		// 		if err != nil {
		// 			continue
		// 		}
		// 		userDataUserID := userData["user"].(map[string]interface{})["id"].(int)
		// 		userDataUserBeforeMemberID := userData["user_before_member"].(map[string]interface{})["id"].(uint32)
		// 		if userDataUserBeforeMemberID <= 0 {
		// 			continue
		// 		}

		// 		belongStudioID := userData["belong_studio_id"].(uint32)
		// 		if belongStudioID <= 0 {
		// 			continue
		// 		}

		// 		strategyType := h.matchUserStrategyType(c, userData)
		// 		studioStrategyID := studioStrategyTypeMap[belongStudioID][strategyType]
		// 		if studioStrategyID <= 0 {
		// 			continue
		// 		}

		// 		strategyIndicatorsData, ok := strategyIndicatorsDataMap[studioStrategyID]
		// 		if !ok {
		// 			continue
		// 		}

		// 		var totalScore float64
		// 		for _, strategyIndicator := range strategyIndicatorsData.StrategyIndicators {
		// 			score := strategyIndicatorCalculateFunc(&strategyIndicator, userData)
		// 			if score == nil || score.ID <= 0 {
		// 				continue
		// 			}
		// 			totalScore += score.Score
		// 		}

		// 		redisKey := fmt.Sprintf(
		// 			"%s_%s_%s",
		// 			strategyRecommendIDs,
		// 			strconv.FormatUint(uint64(belongStudioID), 10),
		// 			strconv.FormatUint(uint64(strategyType), 10),
		// 		)

		// 		h.redis.ZAdd(
		// 			context.TODO(),
		// 			redisKey,
		// 			&redis.Z{
		// 				Score:  totalScore,
		// 				Member: fmt.Sprintf("%v_%v", userDataUserID, userDataUserBeforeMemberID),
		// 			},
		// 		)
		// 		zcard := h.redis.ZCard(context.TODO(), redisKey).Val()
		// 		if zcard > 1000 {
		// 			h.redis.ZRemRangeByRank(context.TODO(), redisKey, 0, zcard-1000)
		// 		}
		// 	}
		// 	wg.Done()
		// }(userBeforeMemberIDs)

	}
	wg.Wait()

	return nil
}
