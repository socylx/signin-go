package v2

import (
	"gsteps-go/api/v2/strategy"
	"gsteps-go/api/v2/strategy_indicator"
	"gsteps-go/api/v2/user_snapshot"
	"gsteps-go/internal/core"
	"gsteps-go/repository/permission"
	"gsteps-go/router/middleware"
)

func Router(routerGroup core.RouterGroup) {
	strategy.CheckPermissionRouter(
		routerGroup.Group("/strategy", middleware.CheckPermission(permission.Admin)),
	)
	strategy.UnCheckPermissionRouter(
		routerGroup.Group("/strategy"),
	)
	strategy_indicator.Router(
		routerGroup.Group("/strategy_indicator", middleware.CheckPermission(permission.Admin)),
	)
	user_snapshot.Router(
		routerGroup.Group("/user_snapshot", middleware.CheckPermission(permission.Admin)),
	)
}
