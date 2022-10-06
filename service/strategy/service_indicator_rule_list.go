package strategy

import (
	"signin-go/internal/core"
	"signin-go/repository/strategy"
	"signin-go/repository/strategy_indicator"
	"signin-go/repository/strategy_indicator_rule_map"
)

func IndicatorDataList(ctx core.StdContext, strategyID uint32) (datas []*strategy.IndicatorData, err error) {
	strategyIndicatorRuleIDs := []uint32{}
	scoreMap := map[uint32]uint32{}
	ruleWeightMap := map[uint32]uint32{}

	strategyIndicatorRuleMaps, _ := strategy_indicator_rule_map.GetStrategyIndicatorRuleMaps(ctx, strategyID)
	for _, strategyIndicatorRuleMap := range strategyIndicatorRuleMaps {
		strategyIndicatorRuleIDs = append(strategyIndicatorRuleIDs, strategyIndicatorRuleMap.StrategyIndicatorRuleID)
		ruleWeightMap[strategyIndicatorRuleMap.StrategyIndicatorRuleID] = strategyIndicatorRuleMap.Weight
		scoreMap[strategyIndicatorRuleMap.StrategyIndicatorRuleID] = strategyIndicatorRuleMap.Score
	}

	weightMap := map[uint32]uint32{}
	strategyIndicatorIDs := []uint32{}
	strategyIndicatorRuleDatasMap := map[uint32][]*strategy.IndicatorRuleData{}

	strategyIndicatorRuleDatas, _ := strategy.GetStrategyIndicatorRuleDatas(ctx, strategyIndicatorRuleIDs)
	for _, strategyIndicatorRuleData := range strategyIndicatorRuleDatas {
		strategyIndicatorIDs = append(strategyIndicatorIDs, strategyIndicatorRuleData.StrategyIndicatorID)
		strategyIndicatorRuleData.Score = scoreMap[strategyIndicatorRuleData.ID]
		strategyIndicatorRuleDatasMap[strategyIndicatorRuleData.StrategyIndicatorID] = append(strategyIndicatorRuleDatasMap[strategyIndicatorRuleData.StrategyIndicatorID], strategyIndicatorRuleData)
		weightMap[strategyIndicatorRuleData.StrategyIndicatorID] = ruleWeightMap[strategyIndicatorRuleData.ID]
	}

	strategyIndicators, _ := strategy_indicator.List(ctx, strategyIndicatorIDs)
	for _, strategyIndicator := range strategyIndicators {
		data := &strategy.IndicatorData{
			ID:     strategyIndicator.ID,
			Key:    strategyIndicator.Key,
			Name:   strategyIndicator.Name,
			Weight: weightMap[strategyIndicator.ID],
			Rules:  strategyIndicatorRuleDatasMap[strategyIndicator.ID],
		}
		datas = append(datas, data)
	}
	return
}
