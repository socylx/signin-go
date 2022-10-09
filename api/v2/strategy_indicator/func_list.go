package strategy_indicator

import (
	"gsteps-go/internal/code"
	"gsteps-go/internal/core"
	"gsteps-go/repository/strategy_indicator"
	"gsteps-go/repository/strategy_indicator_rule"
)

type strategyIndicatorData struct {
	strategy_indicator.StrategyIndicator
	StrategyIndicatorRules []*strategy_indicator_rule.StrategyIndicatorRule `json:"rules"`
}

func list(c core.Context) {
	strategyIndicators, err := strategy_indicator.List(c.RequestContext(), []uint32{})
	if err != nil {
		c.AbortWithError(core.Error(
			code.StrategyIndicatorQueryError,
			code.Text(code.StrategyIndicatorQueryError)).WithError(err),
		)
		return
	}
	strategyIndicatorDatas := make([]*strategyIndicatorData, 0, len(strategyIndicators))
	for _, strategyIndicator := range strategyIndicators {
		strategyIndicatorRules, _ := strategy_indicator_rule.GetStrategyIndicatorRules(c.RequestContext(), []uint32{strategyIndicator.ID})

		strategyIndicatorData := &strategyIndicatorData{}
		strategyIndicatorData.ID = strategyIndicator.ID
		strategyIndicatorData.Name = strategyIndicator.Name
		strategyIndicatorData.StrategyIndicatorRules = strategyIndicatorRules

		strategyIndicatorDatas = append(strategyIndicatorDatas, strategyIndicatorData)
	}
	c.Payload(&strategyIndicatorDatas)
}
