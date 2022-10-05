package strategy

import (
	"fmt"
	redisGlob "signin-go/global/redis"
	"signin-go/internal/code"
	"signin-go/internal/core"
	strategyRepo "signin-go/repository/strategy"
	studioRepo "signin-go/repository/studio"
	"signin-go/repository/user_before_member"
	"signin-go/service/strategy_indicator"
	studioServ "signin-go/service/studio"
	"signin-go/service/users"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v8"
)

const strategyRecommendIDs string = "StrategyRecommendIDs"

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

		wg.Add(1)
		go func(userBeforeMemberIDs []int64) {
			for _, userBeforeMemberID := range userBeforeMemberIDs {
				userData, err := users.Data(ctx, &users.DataID{
					UserID:             0,
					UserBeforeMemberID: uint32(userBeforeMemberID),
				})
				if err != nil {
					continue
				}
				userID := userData.User.ID
				userBeforeMemberID := userData.UserBeforeMember.ID
				if userBeforeMemberID <= 0 {
					continue
				}
				var belongStudioID uint32
				if userData.User.BelongsStudioID > 0 {
					belongStudioID = userData.User.BelongsStudioID
				} else if userData.UserBeforeMember.BelongStudioID > 0 {
					belongStudioID = userData.UserBeforeMember.BelongStudioID
				}
				if belongStudioID <= 0 {
					continue
				}
				strategyType := MatchUserStrategyType(ctx, userData)
				studioStrategyID := studioStrategyTypeIDMap[belongStudioID][strategyType]
				if studioStrategyID <= 0 {
					continue
				}

				strategyIndicatorsData, ok := strategyIndicatorsDataMap[studioStrategyID]
				if !ok {
					continue
				}

				var totalScore float64
				for _, strategyIndicator := range strategyIndicatorsData.StrategyIndicators {
					score, err := strategy_indicator.StrategyIndicatorCalculate(strategyIndicator, userData)
					if err != nil {
						continue
					}
					totalScore += score.Score
				}
				redisKey := fmt.Sprintf("%s_%s_%s", strategyRecommendIDs, strconv.FormatUint(uint64(belongStudioID), 10), strconv.FormatUint(uint64(strategyType), 10))
				redisGlob.Redis.ZAdd(
					ctx,
					redisKey,
					&redis.Z{
						Score:  totalScore,
						Member: fmt.Sprintf("%v_%v", userID, userBeforeMemberID),
					},
				)
				zcard, err := redisGlob.Redis.ZCard(ctx, redisKey).Result()
				if err == nil && zcard > 1000 {
					redisGlob.Redis.ZRemRangeByRank(ctx, redisKey, 0, zcard-1000)
				}
			}
			wg.Done()
		}(userBeforeMemberIDs)
		lastUserBeforeMemberID += 1000
	}
	wg.Wait()

	return nil
}
