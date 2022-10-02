package users

import (
	"signin-go/internal/core"
	"signin-go/service/users"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	detail(c core.Context)
	update(c core.Context)
}

type handler struct {
	userService users.Service
}

func Router(routerGroup core.RouterGroup) {
	h := &handler{
		userService: users.New(),
	}
	rg := routerGroup.Group("/users")
	{
		rg.GET("/detail", h.detail)
		rg.GET("/update", h.update)
		rg.GET("/list", h.list)
	}
}
