package users

import (
	"errors"
	"log"
	"signin-go/global/logger"
	"signin-go/internal/core"

	"go.uber.org/zap"
)

func (h *handler) detail(c core.Context) {
	log.Println("detail")

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
	// detail, _ := h.userService.List(c, 3352)

	log.Println("sssssss.detail")

	c.Payload("detail")
}
