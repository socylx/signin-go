package strategy_indicator

import "gsteps-go/internal/core"

func Router(routerGroup core.RouterGroup) {
	routerGroup.POST("/list", list)
}
