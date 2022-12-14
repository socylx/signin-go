package strategy

import (
	"gsteps-go/global/mysql"
	"gsteps-go/global/time"
	"gsteps-go/global/utils"
	"gsteps-go/internal/code"
	"gsteps-go/internal/core"
	"gsteps-go/internal/validation"
	strategyRepo "gsteps-go/repository/strategy"
	"gsteps-go/repository/strategy_indicator_rule_map"
	"gsteps-go/repository/studio_strategy_map"
	"gsteps-go/service/strategy"
)

type addRequest struct {
	Name                  string `form:"name" binding:"required"`
	Desc                  string `form:"desc" binding:"required"`
	Type                  uint32 `form:"type" binding:"required"`
	StrategyIndicatorRule string `form:"strategy_indicator_rule" binding:"required"`
	StudioIDs             string `form:"studio_ids"`
}

type strategyIndicatorRuleParams struct {
	ID     uint32 `json:"id"`
	Weight uint   `json:"weight"`
	Score  uint   `json:"score"`
}

func add(c core.Context) {
	request := new(addRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}

	valid := strategyRepo.StrategyTypeCheck[request.Type]
	if !valid {
		c.AbortWithError(core.Error(code.ParamBindError, "策略类型不支持"))
		return
	}

	var strategyIndicatorRule []*strategyIndicatorRuleParams
	if err := utils.Json.Unmarshal([]byte(request.StrategyIndicatorRule), &strategyIndicatorRule); err != nil {
		c.AbortWithError(core.Error(code.ParamBindError, "指标规则数据不规范").WithError(err))
		return
	}

	var studioIDs []uint32
	if request.StudioIDs != "" {
		if err := utils.Json.Unmarshal([]byte(request.StudioIDs), &studioIDs); err != nil {
			c.AbortWithError(core.Error(code.ParamBindError, "门店参数不规范").WithError(err))
			return
		}
	}

	s := strategyRepo.Strategy{
		Name:         request.Name,
		Desc:         request.Desc,
		Type:         request.Type,
		Key:          strategy.GenerateStrategyKey(),
		CreateUserID: uint32(c.SessionUserInfo().UserID),
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}

	db := mysql.DB.WithContext(c.RequestContext())

	err := db.Create(&s).Error
	if err != nil {
		c.AbortWithError(core.Error(code.StrategyCreateError, code.Text(code.StrategyCreateError)).WithError(err))
		return
	}

	var strategyIndicatorRuleMaps = []*strategy_indicator_rule_map.StrategyIndicatorRuleMap{}
	for _, sir := range strategyIndicatorRule {
		strategyIndicatorRuleMaps = append(
			strategyIndicatorRuleMaps,
			&strategy_indicator_rule_map.StrategyIndicatorRuleMap{
				StrategyID:              s.ID,
				StrategyIndicatorRuleID: sir.ID,
				Weight:                  uint32(sir.Weight),
				Score:                   uint32(sir.Score),
				CreateTime:              time.Now(),
				UpdateTime:              time.Now(),
			},
		)
	}
	db.Create(&strategyIndicatorRuleMaps)

	var studioStrategyMaps = []*studio_strategy_map.StudioStrategyMap{}
	for _, studioID := range studioIDs {
		studioStrategyMaps = append(
			studioStrategyMaps,
			&studio_strategy_map.StudioStrategyMap{
				StudioID:   studioID,
				StrategyID: s.ID,
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			},
		)
	}
	db.Create(&studioStrategyMaps)

	c.Payload("success")
}
