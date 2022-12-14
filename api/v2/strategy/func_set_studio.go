package strategy

import (
	"gsteps-go/global/time"
	"gsteps-go/global/utils"
	"gsteps-go/internal/code"
	"gsteps-go/internal/core"
	"gsteps-go/internal/validation"
	strategyRepo "gsteps-go/repository/strategy"
	"gsteps-go/repository/studio_strategy_map"
	strategyServ "gsteps-go/service/strategy"
	studioServ "gsteps-go/service/studio"
)

type setStudioRequest struct {
	ID        uint32 `form:"id" binding:"required"`
	StudioIDs string `form:"studio_ids"`
}

func setStudio(c core.Context) {
	request := new(setStudioRequest)
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

	var studioIDs []uint32
	if request.StudioIDs != "" {
		if err := utils.Json.Unmarshal([]byte(request.StudioIDs), &studioIDs); err != nil {
			c.AbortWithError(core.Error(code.ParamBindError, "门店参数不规范").WithError(err))
			return
		}
	}

	strategyDetail, _ := strategyRepo.Detail(c.RequestContext(), request.ID)
	if strategyDetail.ID <= 0 {
		c.AbortWithError(core.Error(code.StrategyQueryError, "无对应策略"))
		return
	}

	studio_strategy_map.Delete(c.RequestContext(), request.ID)
	strategyRepo.Update(c.RequestContext(), request.ID, map[string]interface{}{
		"strategy.key": strategyServ.GenerateStrategyKey(),
	})

	if len(studioIDs) > 0 {
		studioStrategyTypeIDMap, err := studioServ.GetStudioStrategyTypeIDMap(c.RequestContext(), studioIDs)
		if err != nil {
			c.AbortWithError(core.Error(
				code.StudioQueryError,
				"获取门店应用的某种类型的策略对应的策略ID失败").WithError(err),
			)
			return
		}

		var studioStrategyMaps = []*studio_strategy_map.StudioStrategyMap{}
		for _, studioID := range studioIDs {
			if strategyDetail.Status && studioStrategyTypeIDMap[studioID][strategyDetail.Type] > 0 && studioStrategyTypeIDMap[studioID][strategyDetail.Type] != strategyDetail.ID {
				continue
			}
			studioStrategyMaps = append(
				studioStrategyMaps,
				&studio_strategy_map.StudioStrategyMap{
					StudioID:   studioID,
					StrategyID: request.ID,
					CreateTime: time.Now(),
					UpdateTime: time.Now(),
				},
			)
		}
		if len(studioStrategyMaps) > 0 {
			err := studio_strategy_map.Creates(c.RequestContext(), studioStrategyMaps)
			if err != nil {
				c.AbortWithError(core.Error(
					code.StudioStrategyMapCreateError,
					code.Text(code.StudioStrategyMapCreateError)).WithError(err),
				)
				return
			}
		}
	}
	c.Payload("success")
}
