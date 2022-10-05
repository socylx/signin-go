package strategy

import (
	"signin-go/internal/core"
	"signin-go/repository/strategy"
)

func Types(c core.Context) {
	c.Payload(strategy.StrategyTypes)
}
