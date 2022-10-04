package v2

import (
	"signin-go/api/v2/strategy"
	"signin-go/api/v2/strategy_indicator"
	"signin-go/api/v2/user_snapshot"
	"signin-go/internal/core"
	"signin-go/repository/permission"
	"signin-go/router/middleware"
)

func Router(routerGroup core.RouterGroup) {
	strategy.Router(
		routerGroup.Group("/strategy", middleware.CheckPermission(permission.Admin)),
	)
	strategy_indicator.Router(
		routerGroup.Group("/strategy_indicator", middleware.CheckPermission(permission.Admin)),
	)
	user_snapshot.Router(
		routerGroup.Group("/user_snapshot", middleware.CheckPermission(permission.Admin)),
	)
}
