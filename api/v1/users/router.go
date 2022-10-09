package users

import (
	"gsteps-go/internal/core"
)

func Router(routerGroup core.RouterGroup) {
	routerGroup.GET("/detail", detail)
	routerGroup.GET("/update", update)
	routerGroup.GET("/list", list)
}
