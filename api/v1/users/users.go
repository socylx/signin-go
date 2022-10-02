package users

import "signin-go/internal/core"

func Init(routerGroup core.RouterGroup) {
	rg := routerGroup.Group("/users")
	{
		rg.GET("/detail", detail)
	}
}
