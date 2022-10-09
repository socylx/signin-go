package strategy

import (
	"gsteps-go/internal/code"
	"gsteps-go/internal/core"
	"gsteps-go/internal/validation"
	"gsteps-go/service/strategy"
)

type indicatorRuleListRequest struct {
	ID uint32 `form:"id" binding:"required"`
}

func indicatorRuleList(c core.Context) {
	request := new(indicatorRuleListRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}
	strategyIndicators, _ := strategy.IndicatorDataList(c.RequestContext(), request.ID)
	c.Payload(strategyIndicators)
}
