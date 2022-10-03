package users

import (
	"signin-go/global/logger"
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/internal/validation"
	"signin-go/repository/users"

	"go.uber.org/zap"
)

type listRequest struct {
	ID uint32 `form:"id" binding:"required"`
}

func list(c core.Context) {
	request := new(listRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}

	list, err := users.List(c.RequestContext(), &users.Filter{ID: request.ID})
	if err != nil {
		c.AbortWithError(core.Error(
			code.UsersDetailError,
			code.Text(code.UsersDetailError)).WithError(err),
		)
		logger.Logger.Error(
			"users.list",
			zap.Uint32("ID", request.ID),
		)
		return
	}

	c.Payload(list)
}
