package strategy

import (
	"signin-go/global/time"
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/internal/validation"
	"signin-go/repository/strategy"
	"signin-go/repository/studio_strategy_map"
	"signin-go/service/studio"
)

type setStatusRequest struct {
	ID     uint32 `form:"id" binding:"required"`
	Status uint   `form:"status"`
}

func setStatus(c core.Context) {
	request := new(setStatusRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}
	if request.ID <= 0 {
		c.AbortWithError(core.Error(code.ParamBindError, "无id参数"))
		return
	}
	strategyDetail, _ := strategy.Detail(c.RequestContext(), request.ID)
	if strategyDetail.ID <= 0 {
		c.AbortWithError(core.Error(code.StrategyQueryError, "无对应策略"))
		return
	}

	var updateData = map[string]interface{}{
		"strategy.status": request.Status,
	}
	if request.Status == 1 {
		updateData["strategy.start_time"] = time.Now()
		studioIDs, err := studio_strategy_map.GetStudioIDs(c.RequestContext(), request.ID)
		if err != nil {
			c.AbortWithError(core.Error(code.StudioQueryError, "查询此策略应用的门店失败").WithError(err))
			return
		}
		if len(studioIDs) > 0 {
			studioStrategyTypeIDMap, err := studio.GetStudioStrategyTypeIDMap(c.RequestContext(), studioIDs)
			if err != nil {
				c.AbortWithError(core.Error(code.StudioQueryError, "获取门店应用的策略对应的策略ID失败").WithError(err))
				return
			}
			for _, studioID := range studioIDs {
				if studioStrategyTypeIDMap[studioID][strategyDetail.Type] > 0 {
					c.AbortWithError(core.Error(code.SetStrategyStatusError, "已有门店设置相同的类型的策略"))
					return
				}
			}
		}
	}
	err := strategy.Update(c.RequestContext(), request.ID, updateData)
	if err != nil {
		c.AbortWithError(core.Error(
			code.StrategyUpdateError,
			code.Text(code.StrategyUpdateError)).WithError(err),
		)
		return
	}
	c.Payload("success")
}
