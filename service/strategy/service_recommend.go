package strategy

import (
	"fmt"
	redisGlob "signin-go/global/redis"
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/repository/membership"
	redisRepo "signin-go/repository/redis"
	"signin-go/repository/strategy"
	studioRepo "signin-go/repository/studio"
	"signin-go/repository/user_before_member"
	"signin-go/service/strategy_indicator"
	studioServ "signin-go/service/studio"
	"signin-go/service/users"
	"sync"

	"github.com/go-redis/redis/v8"
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

	strategyIndicatorsDataMap := map[uint32]*strategy.StrategyDocument{}
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
		if lastUserBeforeMemberID > maxUserBeforeMemberID || lastUserBeforeMemberID >= 1000 {
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
				redisKey := redisRepo.GetStrategyRecommendIDsRedisKey(belongStudioID, strategyType)
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
	}
	wg.Wait()

	return nil
}

func GenerateOfRenew(ctx core.StdContext) core.BusinessError {
	studioIDs, err := studioRepo.GetStudioIDs(ctx)
	if err != nil || len(studioIDs) <= 0 {
		return core.Error(code.StudioQueryError, code.Text(code.StudioQueryError)).WithError(err)
	}

	studioStrategyTypeIDMap, err := studioServ.GetStudioStrategyTypeIDMap(ctx, studioIDs)
	if err != nil {
		return core.Error(code.StudioQueryError, "获取门店应用的某种类型的策略对应的策略ID").WithError(err)
	}

	strategyIndicatorsDataMap := map[uint32]*strategy.StrategyDocument{}
	for _, strategyTypeID := range studioStrategyTypeIDMap {
		for _, strategyID := range strategyTypeID {
			strategyIndicatorsData, err := Data(ctx, strategyID)
			if err != nil {
				continue
			}
			strategyIndicatorsDataMap[strategyID] = strategyIndicatorsData
		}
	}

	maxMembershipID, err := membership.GetMaxMembershipID(ctx)
	if err != nil {
		return core.Error(code.UserBeforeMemberQueryError, "查询会员卡ID最大值失败").WithError(err)
	}

	var wg sync.WaitGroup
	var lastMembershipID int64
	for {
		if lastMembershipID > maxMembershipID || lastMembershipID >= 1000 {
			break
		}
		membershipUserIDs, err := membership.GetUserBeforeMemberIDs(ctx, &membership.Filter{IDGT: lastMembershipID})
		lastMembershipID += 1000
		if err != nil || len(membershipUserIDs) <= 0 {
			continue
		}

		wg.Add(1)
		go func(membershipUserIDs []int64) {
			for _, membershipUserID := range membershipUserIDs {
				userData, err := users.Data(ctx, &users.DataID{
					UserID:             uint32(membershipUserID),
					UserBeforeMemberID: 0,
				})
				if err != nil {
					continue
				}
				userID := userData.User.ID
				userBeforeMemberID := userData.UserBeforeMember.ID
				if userID <= 0 {
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
				if strategyType != strategy.Xuka {
					continue
				}
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

				redisKey := redisRepo.GetStrategyRecommendIDsRedisKey(belongStudioID, strategyType)
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
		}(membershipUserIDs)
	}
	wg.Wait()

	return nil
}
