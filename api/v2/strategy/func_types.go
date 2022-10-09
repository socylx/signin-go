package strategy

import (
	"gsteps-go/internal/core"
	"gsteps-go/repository/strategy"
)

func Types(c core.Context) {
	c.Payload(strategy.StrategyTypes)
}
