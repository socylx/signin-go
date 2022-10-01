package router

import (
	"errors"
	"net/http"
	"signin-go/global/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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
		c.String(http.StatusOK, "Gin Server")
	})
	return engine
}
