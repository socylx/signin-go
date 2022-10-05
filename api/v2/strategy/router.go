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
	routerGroup.POST("/recommend", recommend)
	routerGroup.POST("/score", score)
	routerGroup.POST("/scores", scores)
	routerGroup.POST("/set/status", setStatus)

}

func UnCheckPermissionRouter(routerGroup core.RouterGroup) {
	routerGroup.POST("/recommend/generatelaxin", recommendGenerateOfLaxin)
	routerGroup.POST("/recommend/generaterenew", recommendGenerateOfRenew)
}
