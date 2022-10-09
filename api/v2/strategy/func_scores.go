package strategy

import (
	"gsteps-go/global/mongo"
	"gsteps-go/global/time"
	"gsteps-go/internal/code"
	"gsteps-go/internal/core"
	"gsteps-go/internal/validation"
	strategyRepo "gsteps-go/repository/strategy"
	"gsteps-go/repository/users"
	strategyServ "gsteps-go/service/strategy"
	"gsteps-go/service/strategy_indicator"
	studioServ "gsteps-go/service/studio"
	usersServ "gsteps-go/service/users"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

type scoresRequest struct {
	UserID             uint32 `form:"user_id"`
	UserBeforeMemberID uint32 `form:"user_before_member_id"`
}

func scores(c core.Context) {
	request := new(scoresRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}
	if request.UserID <= 0 && request.UserBeforeMemberID <= 0 {
		c.AbortWithError(core.Error(code.ParamBindError, "无id参数"))
		return
	}

	userData, err := usersServ.Data(c.RequestContext(), &usersServ.DataID{
		UserID:             request.UserID,
		UserBeforeMemberID: request.UserBeforeMemberID,
	})
	if err != nil {
		c.AbortWithError(core.Error(
			code.UsersDataError,
			code.Text(code.UsersDataError)).WithError(err),
		)
		return
	}

	var belongStudioID uint32
	if userData.User.BelongsStudioID > 0 {
		belongStudioID = userData.User.BelongsStudioID
	} else if userData.UserBeforeMember.BelongStudioID > 0 {
		belongStudioID = userData.UserBeforeMember.BelongStudioID
	}
	if belongStudioID <= 0 {
		c.AbortWithError(core.Error(code.UsersDataError, "该用户无所属门店"))
		return
	}

	studioStrategyTypeIDMap, err := studioServ.GetStudioStrategyTypeIDMap(c.RequestContext(), []uint32{belongStudioID})
	if err != nil {
		c.AbortWithError(core.Error(
			code.StudioQueryError,
			"获取门店应用的某种类型的策略对应的策略ID").WithError(err),
		)
		return
	}

	userID := userData.User.ID
	var strategyIndicatorScores []*users.StrategyIndicatorScore
	strategyIndicatorScoreType := map[uint32]bool{}
	coll := mongo.Mongo.Collection(users.StrategyIndicatorScoreName)
	cursor, err := coll.Find(c.RequestContext(), bson.D{{Key: "user_id", Value: userID}})
	if err == nil {
		for cursor.Next(c.RequestContext()) {
			var elem users.StrategyIndicatorScore
			err := cursor.Decode(&elem)
			if err == nil {
				strategyIndicatorScoreType[elem.Type] = true
				strategyIndicatorScores = append(strategyIndicatorScores, &elem)
			}
		}
		if err := cursor.Err(); err == nil {
			cursor.Close(c.RequestContext())
		}
	}

	var wg sync.WaitGroup
	mu := sync.Mutex{}
	appendScore := func(score *users.StrategyIndicatorScore) {
		mu.Lock()
		defer func() {
			mu.Unlock()
		}()
		strategyIndicatorScores = append(strategyIndicatorScores, score)
	}
	for k := range strategyRepo.StrategyTypeCheck {
		if strategyIndicatorScoreType[k] {
			continue
		}

		wg.Add(1)
		go func(strategyType uint32) {
			defer func() {
				wg.Done()
			}()

			studioStrategyID := studioStrategyTypeIDMap[belongStudioID][strategyType]
			if studioStrategyID <= 0 {
				return
			}

			strategiesData, err := strategyRepo.List(c.RequestContext(), &strategyRepo.ListFilter{
				IncludeIds: []uint32{studioStrategyID},
				Status:     1,
				Type:       strategyType,
				Page:       1,
				Size:       2,
			})
			if err != nil {
				return
			}
			if len(strategiesData.Data) <= 0 {
				return
			}
			strategyData := strategiesData.Data[0]
			strategyIndicatorsData, err := strategyServ.Data(c.RequestContext(), strategyData.ID)
			if err != nil {
				return
			}

			var scores []*users.Score
			for _, strategyIndicator := range strategyIndicatorsData.StrategyIndicators {
				score, err := strategy_indicator.StrategyIndicatorCalculate(strategyIndicator, userData)
				if err != nil {
					continue
				}
				scores = append(scores, score)
			}
			if len(scores) > 0 {
				appendScore(&users.StrategyIndicatorScore{
					UserID:      uint32(userID),
					Time:        time.Now(),
					Type:        strategyType,
					StrategyKey: strategyData.Key,
					// Data:        userData,
					Scores: scores,
				})
			}
		}(k)
	}
	wg.Wait()

	c.Payload(strategyIndicatorScores)
}
