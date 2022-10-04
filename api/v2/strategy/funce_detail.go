package strategy

import (
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/internal/validation"
	"signin-go/repository/strategy"
)

type detailRequest struct {
	ID uint32 `form:"id" binding:"required"`
}

func detail(c core.Context) {
	request := new(detailRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}

	res, err := strategy.Detail(c.RequestContext(), request.ID)
	if err != nil {
		c.AbortWithError(core.Error(code.StrategyQueryError, code.Text(code.StrategyQueryError)).WithError(err))
		return
	}
	c.Payload(res)
}
