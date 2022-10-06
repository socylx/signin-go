package strategy

import (
	"signin-go/global/mongo"
	"signin-go/global/mysql"
	"signin-go/global/utils"
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/internal/validation"
	strategyRepo "signin-go/repository/strategy"
	"signin-go/repository/strategy_indicator_rule_map"
	"signin-go/repository/studio_strategy_map"
	strategyServ "signin-go/service/strategy"

	"go.mongodb.org/mongo-driver/bson"
)

type updateRequest struct {
	ID                    uint32 `form:"id" binding:"required"`
	Name                  string `form:"name" binding:"required"`
	Desc                  string `form:"desc" binding:"required"`
	Type                  uint32 `form:"type" binding:"required"`
	StrategyIndicatorRule string `form:"strategy_indicator_rule" binding:"required"`
	StudioIDs             string `form:"studio_ids"`
}

func update(c core.Context) {
	request := new(updateRequest)
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

	data, err := strategyServ.Data(c.RequestContext(), request.ID)
	if err != nil {
		c.AbortWithError(core.Error(code.StrategyQueryError, "查询策略数据失败").WithError(err))
		return
	}

	strategyColl := mongo.Mongo.Collection(strategyRepo.StrategyName)
	var strategy strategyRepo.StrategyDocument
	if err = strategyColl.FindOne(c.RequestContext(), bson.M{"key": data.Key}).Decode(&strategy); err != nil {
		strategyColl.InsertOne(c.RequestContext(), data)
	}

	err = strategyRepo.Update(c.RequestContext(), request.ID, map[string]interface{}{
		"strategy.name":   request.Name,
		"strategy.desc":   request.Desc,
		"strategy.type":   request.Type,
		"strategy.key":    strategyServ.GenerateStrategyKey(),
		"strategy.status": 0,
	})

	if err != nil {
		c.AbortWithError(core.Error(code.StrategyUpdateError, code.Text(code.StrategyUpdateError)).WithError(err))
		return
	}

	strategy_indicator_rule_map.Delete(c.RequestContext(), request.ID)
	studio_strategy_map.Delete(c.RequestContext(), request.ID)

	var strategyIndicatorRuleMaps = []*strategy_indicator_rule_map.StrategyIndicatorRuleMap{}
	for _, sir := range strategyIndicatorRule {
		strategyIndicatorRuleMaps = append(
			strategyIndicatorRuleMaps,
			&strategy_indicator_rule_map.StrategyIndicatorRuleMap{
				StrategyID:              request.ID,
				StrategyIndicatorRuleID: sir.ID,
				Weight:                  uint32(sir.Weight),
				Score:                   uint32(sir.Score),
			},
		)
	}
	if len(strategyIndicatorRuleMaps) > 0 {
		mysql.DB.Create(&strategyIndicatorRuleMaps)
	}

	var studioStrategyMaps = []*studio_strategy_map.StudioStrategyMap{}
	for _, studioID := range studioIDs {
		studioStrategyMaps = append(
			studioStrategyMaps,
			&studio_strategy_map.StudioStrategyMap{
				StudioID:   studioID,
				StrategyID: request.ID,
			},
		)
	}
	if len(studioStrategyMaps) > 0 {
		mysql.DB.Create(&studioStrategyMaps)
	}
	c.Payload("success")
}
