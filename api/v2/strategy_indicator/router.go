package strategy_indicator

import "signin-go/internal/core"

func Router(routerGroup core.RouterGroup) {
	routerGroup.POST("/list", list)
}
