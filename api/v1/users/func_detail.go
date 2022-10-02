package users

import (
	"errors"
	"signin-go/global/logger"
	"signin-go/internal/core"

	"go.uber.org/zap"
)

func (h *handler) detail(c core.Context) {
	var err error = errors.New("sss")
	logger.Logger.Error(
		"err occurs",
		zap.Any("method", "sss"),
		zap.Error(err),
		// logger.WrapMeta(
		// 	nil,
		// 	logger.NewMeta("para1", "value1"),
		// 	logger.NewMeta("para2", "value2"),
		// )...,
	)
	list, _ := h.userService.Detail(c, 3352)

	c.Payload(list)
}
