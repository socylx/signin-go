package strategy_indicator

import (
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/repository/strategy_indicator"
)

func list(c core.Context) {
	strategyIndicators, err := strategy_indicator.List(c.RequestContext())
	if err != nil {
		c.AbortWithError(core.Error(
			code.StrategyIndicatorQueryError,
			code.Text(code.StrategyIndicatorQueryError)).WithError(err),
		)
		return
	}
	c.Payload(strategyIndicators)
}
