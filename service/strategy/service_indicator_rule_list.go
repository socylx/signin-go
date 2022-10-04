package strategy

import (
	"signin-go/internal/core"
	"signin-go/repository/strategy"
	"signin-go/repository/strategy_indicator_rule_map"
)

func IndicatorDataList(ctx core.StdContext, strategyID uint32) (strategyIndicatorDatas []*strategy.IndicatorData, err error) {
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

	strategyIndicatorDatas, err = strategy.GetStrategyIndicatorDatas(ctx, strategyIndicatorIDs)
	for _, strategyIndicatorData := range strategyIndicatorDatas {
		strategyIndicatorData.Weight = weightMap[strategyIndicatorData.ID]
		strategyIndicatorData.Rules = strategyIndicatorRuleDatasMap[strategyIndicatorData.ID]
	}
	return
}
