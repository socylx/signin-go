package strategy

import (
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/internal/validation"
	strategyRepo "signin-go/repository/strategy"
	usersRepo "signin-go/repository/users"
	"sync"

	strategyServ "signin-go/service/strategy"
	"signin-go/service/strategy_indicator"
	studioServ "signin-go/service/studio"
	usersServ "signin-go/service/users"
)

type scoreRequest struct {
	UserID             uint32 `form:"user_id"`
	UserBeforeMemberID uint32 `form:"user_before_member_id"`
}

type scoreResponse struct {
	UserID             uint32             `json:"user_id"`
	UserBeforeMemberID uint32             `json:"user_before_member_id"`
	Type               uint32             `json:"type"`
	StrategyKey        string             `json:"strategy_key"`
	Scores             []*usersRepo.Score `json:"scores"`
}

func score(c core.Context) {
	request := new(scoreRequest)
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

	strategyType := strategyServ.MatchUserStrategyType(c.RequestContext(), userData)
	studioStrategyID := studioStrategyTypeIDMap[belongStudioID][strategyType]
	if studioStrategyID <= 0 {
		c.AbortWithError(core.Error(code.StudioQueryError, "用户对应的门店无策略"))
		return
	}

	userID := userData.User.ID
	userBeforeMemberID := userData.UserBeforeMember.ID

	strategiesData, err := strategyRepo.List(c.RequestContext(), &strategyRepo.ListFilter{
		IncludeIds: []uint32{studioStrategyID},
		Status:     1,
		Type:       strategyType,
		Page:       1,
		Size:       2,
	})
	if err != nil {
		c.AbortWithError(core.Error(code.StrategyQueryError, code.Text(code.StrategyQueryError)).WithError(err))
		return
	}
	if len(strategiesData.Data) <= 0 {
		c.AbortWithError(core.Error(code.StrategyQueryError, "用户对应的门店无对应类型的策略"))
		return
	}
	strategyData := strategiesData.Data[0]
	strategyIndicatorsData, err := strategyServ.Data(c.RequestContext(), strategyData.ID)
	if err != nil {
		c.AbortWithError(core.Error(code.StrategyQueryError, "查询策略数据失败").WithError(err))
		return
	}

	var wg sync.WaitGroup
	var scores []*usersRepo.Score
	mu := sync.Mutex{}
	appendScore := func(score *usersRepo.Score, err error) {
		defer wg.Done()
		if err != nil {
			return
		}
		defer mu.Unlock()
		mu.Lock()
		scores = append(scores, score)
	}
	for _, strategyIndicator := range strategyIndicatorsData.StrategyIndicators {
		wg.Add(1)
		go appendScore(strategy_indicator.StrategyIndicatorCalculate(strategyIndicator, userData))
	}

	wg.Wait()
	response := &scoreResponse{
		UserID:             userID,
		UserBeforeMemberID: userBeforeMemberID,
		Type:               strategyData.Type,
		StrategyKey:        strategyData.Key,
		Scores:             scores,
	}
	c.Payload(response)
}
