package router

import (
	"errors"
	"net/http"
	"signin-go/global/logger"
	"signin-go/repository/users"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	Code    int         `json:"code"` // 业务码
	Message string      `json:"msg"`  // 描述信息
	Res     interface{} `json:"res"`
}

func HTTPServer() *gin.Engine {
	engine := gin.New()
	engine.GET("/", func(c *gin.Context) {

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

		c.JSON(
			http.StatusOK,
			Response{
				Code:    0,
				Message: "",
				Res:     detail,
			},
		)

	})
	return engine
}
