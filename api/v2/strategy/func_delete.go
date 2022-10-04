package strategy

import (
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/internal/validation"
	"signin-go/repository/strategy"
)

type deleteRequest struct {
	ID uint32 `form:"id" binding:"required"`
}

func delete(c core.Context) {
	request := new(deleteRequest)
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

	err := strategy.Delete(c.RequestContext(), &strategy.DeleteFileter{ID: request.ID})
	if err != nil {
		c.AbortWithError(core.Error(code.StrategyDeleteError, code.Text(code.StrategyDeleteError)).WithError(err))
		return
	}

	c.Payload("success")
}
