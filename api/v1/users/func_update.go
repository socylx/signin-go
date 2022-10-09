package users

import (
	"gsteps-go/global/logger"
	"gsteps-go/internal/code"
	"gsteps-go/internal/core"
	"gsteps-go/internal/validation"
	"gsteps-go/repository/users"

	"go.uber.org/zap"
)

type updateRequest struct {
	ID uint32 `form:"id" binding:"required"`
}

func update(c core.Context) {
	request := new(updateRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}

	err := users.Update(
		c.RequestContext(),
		&users.Filter{ID: request.ID},
		users.Users{Name: "NewName"},
	)
	if err != nil {
		c.AbortWithError(core.Error(
			code.UsersDetailError,
			code.Text(code.UsersDetailError)).WithError(err),
		)
		logger.Logger.Error(
			"users.update",
			zap.Uint32("ID", request.ID),
		)
		return
	}

	c.Payload("success")
}
