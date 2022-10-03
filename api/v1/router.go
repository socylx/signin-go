package v1

import (
	"signin-go/api/v1/users"
	"signin-go/internal/core"
	"signin-go/router/middleware"
)

func Router(routerGroup core.RouterGroup) {
	users.Router(
		routerGroup.Group("/users", middleware.CheckLogin()),
	)
}
