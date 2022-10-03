package v2

import (
	"signin-go/api/v2/user_snapshot"
	"signin-go/internal/core"
	"signin-go/repository/permission"
	"signin-go/router/middleware"
)

func Router(routerGroup core.RouterGroup) {
	user_snapshot.Router(
		routerGroup.Group("/user_snapshot", middleware.CheckPermission(permission.Admin)),
	)
}
