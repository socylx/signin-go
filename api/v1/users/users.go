package users

import (
	"signin-go/internal/core"
	"signin-go/service/users"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	detail(c core.Context)
	list(c core.Context)
	update(c core.Context)
}

type handler struct {
	userService users.Service
}

func Router(routerGroup core.RouterGroup) {
	h := &handler{
		userService: users.New(),
	}

	{
		routerGroup.GET("/detail", h.detail)
		routerGroup.GET("/update", h.update)
		routerGroup.GET("/list", h.list)
	}
}
