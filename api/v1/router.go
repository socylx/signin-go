package v1

import (
	"signin-go/api/v1/users"
	"signin-go/internal/core"
)

func Init(mux core.Mux) {
	routerGroup := mux.Group("/v1")
	users.Router(routerGroup)
}
