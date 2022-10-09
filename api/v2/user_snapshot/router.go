package user_snapshot

import "gsteps-go/internal/core"

func Router(routerGroup core.RouterGroup) {
	routerGroup.POST("/accesstorenew", accesstorenew)
}
