package v1

import (
	"signin-go/api/v1/users"
	"signin-go/internal/core"
)

func Router(routerGroup core.RouterGroup) {

	usersRG := routerGroup.Group("/users")
	users.Router(usersRG)
}
