package strategy

import (
	"signin-go/global/config"
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/internal/validation"
	studioRepo "signin-go/repository/studio"
	"signin-go/service/strategy"
)

var strategyRecommendIDs string = "StrategyRecommendIDs"

type recommendGenerateRequest struct {
	Token string `form:"token" binding:"required"`
}

func recommendGenerateOfLaxin(c core.Context) {
	request := new(recommendGenerateRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}
	if request.Token != config.Server.Token {
		c.AbortWithError(core.Error(code.PermissionError, code.Text(code.PermissionError)))
		return
	}

	studioIDs, err := studioRepo.GetStudioIDs(c.RequestContext())
	if err != nil || len(studioIDs) <= 0 {
		c.AbortWithError(core.Error(
			code.StudioQueryError,
			code.Text(code.StudioQueryError)).WithError(err),
		)
		return
	}

	generr := strategy.GenerateOfLaxin(c.RequestContext())
	if generr != nil {
		c.AbortWithError(generr)
		return
	}

	// var wg sync.WaitGroup
	// var lastUserBeforeMemberID int64
	// for {
	// 	if lastUserBeforeMemberID > maxUserBeforeMemberID {
	// 		logger.Info("recommendGenerateOfLaxin Break...")
	// 		break
	// 	}
	// 	logger.Infof("recommendGenerateOfLaxin.lastUserBeforeMemberID: %v", lastUserBeforeMemberID)

	// 	userBeforeMemberIDs := []int64{}
	// 	h.mysql.Table(
	// 		"user_before_member",
	// 	).Select(
	// 		"user_before_member.id",
	// 	).Where(
	// 		"user_before_member.id > ?",
	// 		lastUserBeforeMemberID,
	// 	).Limit(1000).Find(&userBeforeMemberIDs)

	// 	if len(userBeforeMemberIDs) > 0 {
	// 		wg.Add(1)
	// 		go func(userBeforeMemberIDs []int64) {
	// 			for _, userBeforeMemberID := range userBeforeMemberIDs {
	// 				userData, err := h.userService.Data(c.RequestContext(), &user.DataFilter{
	// 					UserID:             0,
	// 					UserBeforeMemberID: uint32(userBeforeMemberID),
	// 				})
	// 				if err != nil {
	// 					continue
	// 				}
	// 				userDataUserID := userData["user"].(map[string]interface{})["id"].(int)
	// 				userDataUserBeforeMemberID := userData["user_before_member"].(map[string]interface{})["id"].(uint32)
	// 				if userDataUserBeforeMemberID <= 0 {
	// 					continue
	// 				}

	// 				belongStudioID := userData["belong_studio_id"].(uint32)
	// 				if belongStudioID <= 0 {
	// 					continue
	// 				}

	// 				strategyType := h.matchUserStrategyType(c, userData)
	// 				studioStrategyID := studioStrategyTypeMap[belongStudioID][strategyType]
	// 				if studioStrategyID <= 0 {
	// 					continue
	// 				}

	// 				strategyIndicatorsData, ok := strategyIndicatorsDataMap[studioStrategyID]
	// 				if !ok {
	// 					continue
	// 				}

	// 				var totalScore float64
	// 				for _, strategyIndicator := range strategyIndicatorsData.StrategyIndicators {
	// 					score := strategyIndicatorCalculateFunc(&strategyIndicator, userData)
	// 					if score == nil || score.ID <= 0 {
	// 						continue
	// 					}
	// 					totalScore += score.Score
	// 				}

	// 				redisKey := fmt.Sprintf(
	// 					"%s_%s_%s",
	// 					strategyRecommendIDs,
	// 					strconv.FormatUint(uint64(belongStudioID), 10),
	// 					strconv.FormatUint(uint64(strategyType), 10),
	// 				)

	// 				h.redis.ZAdd(
	// 					context.TODO(),
	// 					redisKey,
	// 					&redis.Z{
	// 						Score:  totalScore,
	// 						Member: fmt.Sprintf("%v_%v", userDataUserID, userDataUserBeforeMemberID),
	// 					},
	// 				)
	// 				zcard := h.redis.ZCard(context.TODO(), redisKey).Val()
	// 				if zcard > 1000 {
	// 					h.redis.ZRemRangeByRank(context.TODO(), redisKey, 0, zcard-1000)
	// 				}
	// 			}
	// 			wg.Done()
	// 		}(userBeforeMemberIDs)
	// 	}

	// 	lastUserBeforeMemberID += 1000
	// }
	// wg.Wait()
	// log.Println("recommendGenerateOfLaxin success...")
	c.Payload("success")
}

// func (h *handler) recommendGenerateOfRenew() core.HandlerFunc {
// 	return func(c core.Context) {
// 		request := new(recommendGenerateRequest)
// 		if err := c.ShouldBindForm(request); err != nil {
// 			c.AbortWithError(core.Error(code.ParamBindError, code.Text(code.ParamBindError)).WithError(err))
// 			return
// 		}
// 		if request.Token != config.AppConfig.Token {
// 			c.AbortWithError(core.Error(code.PermissionError, code.Text(code.PermissionError)))
// 			return
// 		}

// 		studioIDs := []uint32{}
// 		h.mysql.Table(
// 			"studio",
// 		).Select(
// 			"studio.id",
// 		).Where(
// 			"studio.is_del = ?", 0,
// 		).Find(&studioIDs)
// 		studioStrategyTypeMap := h.getStudioStrategyTypeMap(studioIDs)

// 		strategyIndicatorsDataMap := map[uint32]*collection.Strategy{}
// 		for _, strategyTypeM := range studioStrategyTypeMap {
// 			for _, strategyID := range strategyTypeM {
// 				strategyIndicatorsData, err := h.strategyService.Data(strategyID)
// 				if err != nil {
// 					continue
// 				}
// 				strategyIndicatorsDataMap[strategyID] = strategyIndicatorsData
// 			}
// 		}

// 		var maxMembershipID int64
// 		if err := h.mysql.Table(
// 			"membership",
// 		).Select(
// 			"membership.id",
// 		).Count(&maxMembershipID).Error; err != nil {
// 			log.Println("recommendGenerateOfRenew.err: ", err)
// 			c.AbortWithError(core.Error(code.ServerError, code.Text(code.ServerError)).WithError(err))
// 			return
// 		}

// 		var wg sync.WaitGroup
// 		var lastMembershipID int64
// 		for {
// 			if lastMembershipID > maxMembershipID {
// 				logger.Info("recommendGenerateOfRenew Break...")
// 				break
// 			}
// 			logger.Infof("recommendGenerateOfRenew.lastMembershipID: %v", lastMembershipID)

// 			membershipUserIDs := []int64{}
// 			h.mysql.Table(
// 				"membership",
// 			).Select(
// 				"membership.user_id",
// 			).Where(
// 				"membership.id > ?", lastMembershipID,
// 			).Limit(1000).Find(&membershipUserIDs)

// 			if len(membershipUserIDs) > 0 {
// 				wg.Add(1)
// 				go func(membershipUserIDs []int64, sIDM map[uint32]*collection.Strategy) {
// 					for _, membershipUserID := range membershipUserIDs {
// 						userData, err := h.userService.Data(c.RequestContext(), &user.DataFilter{
// 							UserID:             uint32(membershipUserID),
// 							UserBeforeMemberID: 0,
// 						})
// 						if err != nil {
// 							continue
// 						}
// 						userDataUserID := userData["user"].(map[string]interface{})["id"].(int)
// 						userDataUserBeforeMemberID := userData["user_before_member"].(map[string]interface{})["id"].(uint32)
// 						if userDataUserID <= 0 {
// 							continue
// 						}

// 						belongStudioID := userData["belong_studio_id"].(uint32)
// 						if belongStudioID <= 0 {
// 							continue
// 						}
// 						strategyType := h.matchUserStrategyType(c, userData)
// 						if strategyType != strategy.StrategyType.Xuka {
// 							continue
// 						}
// 						studioStrategyID := studioStrategyTypeMap[belongStudioID][strategyType]
// 						if studioStrategyID <= 0 {
// 							continue
// 						}
// 						strategyIndicatorsData, ok := sIDM[studioStrategyID]
// 						if !ok {
// 							continue
// 						}

// 						var totalScore float64
// 						for _, strategyIndicator := range strategyIndicatorsData.StrategyIndicators {
// 							score := strategyIndicatorCalculateFunc(&strategyIndicator, userData)
// 							if score == nil || score.ID <= 0 {
// 								continue
// 							}
// 							totalScore += score.Score
// 						}
// 						redisKey := fmt.Sprintf(
// 							"%s_%s_%s",
// 							strategyRecommendIDs,
// 							strconv.FormatUint(uint64(belongStudioID), 10),
// 							strconv.FormatUint(uint64(strategyType), 10),
// 						)
// 						h.redis.ZAdd(
// 							context.TODO(),
// 							redisKey,
// 							&redis.Z{
// 								Score:  totalScore,
// 								Member: fmt.Sprintf("%v_%v", userDataUserID, userDataUserBeforeMemberID),
// 							},
// 						)
// 						zcard := h.redis.ZCard(context.TODO(), redisKey).Val()
// 						if zcard > 1000 {
// 							h.redis.ZRemRangeByRank(context.TODO(), redisKey, 0, zcard-1000)
// 						}
// 					}
// 					wg.Done()
// 				}(membershipUserIDs, strategyIndicatorsDataMap)
// 			}
// 			lastMembershipID += 1000
// 		}
// 		wg.Wait()
// 		log.Println("recommendGenerateOfRenew success...")
// 		c.Payload("success")
// 	}
// }
