package strategy

import "signin-go/internal/core"

func Router(routerGroup core.RouterGroup) {
	routerGroup.POST("/add", add)
	routerGroup.POST("/delete", delete)
}
