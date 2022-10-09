package v1

import (
	"gsteps-go/api/v1/users"
	"gsteps-go/internal/core"
)

func Router(routerGroup core.RouterGroup) {
	users.Router(
		routerGroup.Group("/users"),
	)
}
