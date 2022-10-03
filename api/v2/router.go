package v2

import (
	"signin-go/api/v2/user_snapshot"
	"signin-go/internal/core"
)

func Router(routerGroup core.RouterGroup) {
	user_snapshot.Router(
		routerGroup.Group("/user_snapshot"),
	)
}
