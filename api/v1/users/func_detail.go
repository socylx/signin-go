package users

import (
	"errors"
	"log"
	"signin-go/global/logger"
	"signin-go/internal/core"
	"signin-go/repository/users"

	"go.uber.org/zap"
)

func detail(c core.Context) {
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

	users := users.New()
	detail, _ := users.Detail(3352)

	c.Payload(detail)
}
