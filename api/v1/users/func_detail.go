package users

import (
	"gsteps-go/global/logger"
	"gsteps-go/internal/code"
	"gsteps-go/internal/core"
	"gsteps-go/internal/validation"
	"gsteps-go/repository/users"

	"go.uber.org/zap"
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

	detail, err := users.Detail(c.RequestContext(), request.ID)
	if err != nil {
		c.AbortWithError(core.Error(
			code.UsersDetailError,
			code.Text(code.UsersDetailError)).WithError(err),
		)
		logger.Logger.Error(
			"users.detail",
			zap.Uint32("ID", request.ID),
		)
		return
	}

	c.Payload(detail)
}
