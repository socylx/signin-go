package strategy

import (
	"gsteps-go/internal/code"
	"gsteps-go/internal/core"
	"gsteps-go/internal/validation"
	"gsteps-go/repository/studio_strategy_map"
)

type studioListRequest struct {
	ID uint32 `form:"id" binding:"required"`
}

func studioList(c core.Context) {
	request := new(studioListRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}
	if request.ID <= 0 {
		c.AbortWithError(core.Error(code.ParamBindError, "无id参数"))
		return
	}

	studioIDs, err := studio_strategy_map.GetStudioIDs(c.RequestContext(), request.ID)
	if err != nil {
		c.AbortWithError(core.Error(code.StudioQueryError, "查询此策略应用的门店失败").WithError(err))
		return
	}
	c.Payload(&studioIDs)
}
