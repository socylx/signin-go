package strategy

import (
	"signin-go/global/time"
	"signin-go/global/utils"
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/internal/validation"
	"signin-go/repository/strategy"
	"signin-go/repository/studio_strategy_map"
)

type listRequest struct {
	StartTime    string `form:"start_time"`
	EndTime      string `form:"end_time"`
	Keyword      string `form:"keyword"`
	Status       int    `form:"status"`
	CreateUserID uint32 `form:"create_user_id"`
	Type         uint32 `form:"type"`
	StudioIDs    string `form:"studio_ids"`
	Page         int    `form:"page"`
	Size         int    `form:"size"`
}

func list(c core.Context) {
	request := new(listRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}
	var studioIDs []uint32
	if request.StudioIDs != "" {
		if err := utils.Json.Unmarshal([]byte(request.StudioIDs), &studioIDs); err != nil {
			c.AbortWithError(core.Error(code.ParamBindError, "门店参数不规范").WithError(err))
			return
		}
	}
	if request.Page <= 0 || request.Size <= 0 {
		request.Page = 1
		request.Size = 20
	}
	var startTime, endTime time.Time
	startTime, _ = time.ParseCSTInLocation(request.StartTime, time.YYYYMMDD)
	endTime, _ = time.ParseCSTInLocation(request.EndTime, time.YYYYMMDD)

	var strategyIDs []uint32
	if len(studioIDs) > 0 {
		strategyIDs, _ = studio_strategy_map.GetStudioStrategyIDs(c.RequestContext(), studioIDs)
	}

	list, err := strategy.List(c.RequestContext(), &strategy.ListFilter{
		IncludeIds:   strategyIDs,
		CreateTimeGE: startTime,
		CreateTimeLT: endTime,
		Keyword:      request.Keyword,
		Status:       request.Status,
		CreateUserID: request.CreateUserID,
		Type:         request.Type,
		Page:         request.Page,
		Size:         request.Size,
	})

	if err != nil {
		c.AbortWithError(core.Error(code.StrategyQueryError, code.Text(code.StrategyQueryError)).WithError(err))
		return
	}

	c.Payload(list)
}
