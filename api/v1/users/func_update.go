package users

import (
	"signin-go/global/logger"
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/internal/validation"

	"go.uber.org/zap"
)

type updateRequest struct {
	ID uint32 `form:"id" binding:"required"`
}

func (h *handler) update(c core.Context) {
	request := new(updateRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}

	err := h.userService.Update(c, request.ID)
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
