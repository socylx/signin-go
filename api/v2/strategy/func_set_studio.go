package strategy

import (
	"log"
	"signin-go/global/mysql"
	"signin-go/global/utils"
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/internal/validation"
	strategyRepo "signin-go/repository/strategy"
	"signin-go/repository/studio_strategy_map"
	strategyServ "signin-go/service/strategy"
	studioServ "signin-go/service/studio"
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

	log.Println("studioIDs: ", studioIDs)

	if len(studioIDs) > 0 {
		studioStrategyTypeIDMap, err := studioServ.GetStudioStrategyTypeIDMap(c.RequestContext(), studioIDs)
		if err != nil {
			c.AbortWithError(core.Error(
				code.StudioQueryError,
				"获取门店应用的某种类型的策略对应的策略ID失败").WithError(err),
			)
			return
		}

		var sutidoStrategyMaps = []*studio_strategy_map.StudioStrategyMap{}
		for _, studioID := range studioIDs {
			if strategyDetail.Status && studioStrategyTypeIDMap[studioID][strategyDetail.Type] > 0 && studioStrategyTypeIDMap[studioID][strategyDetail.Type] != strategyDetail.ID {
				log.Println("studioID: ", studioID, " continue")
				continue
			}
			log.Println("studioID: ", studioID, " no continue")
			sutidoStrategyMaps = append(
				sutidoStrategyMaps,
				&studio_strategy_map.StudioStrategyMap{
					StudioID:   studioID,
					StrategyID: request.ID,
				},
			)
		}
		if len(sutidoStrategyMaps) > 0 {
			log.Println("Create: ")
			mysql.DB.WithContext(c.RequestContext()).Create(&sutidoStrategyMaps)
		}
	}
	c.Payload("success")
}
