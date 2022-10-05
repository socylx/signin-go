package strategy

import (
	"signin-go/internal/core"
)

func CheckPermissionRouter(routerGroup core.RouterGroup) {

	routerGroup.POST("/add", add)
	routerGroup.POST("/delete", delete)
	routerGroup.POST("/detail", detail)
	routerGroup.POST("/list", list)
	routerGroup.POST("/indicator/rule/list", indicatorRuleList)

}

func UnCheckPermissionRouter(routerGroup core.RouterGroup) {
	routerGroup.POST("/recommend/generatelaxin", recommendGenerateOfLaxin)
}
