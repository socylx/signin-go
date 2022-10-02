package users

import (
	"signin-go/global/logger"
	"signin-go/internal/core"

	"go.uber.org/zap"
)

func (h *handler) list(c core.Context) {
	list, err := h.userService.List(c, 3352)

	logger.Logger.Error(
		"users.list",
		zap.Any("method", "sss"),
		zap.Error(err),
		// logger.WrapMeta(
		// 	nil,
		// 	logger.NewMeta("para1", "value1"),
		// 	logger.NewMeta("para2", "value2"),
		// )...,
	)
	c.Payload(list)
}
