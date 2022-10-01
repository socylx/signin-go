package router

import (
	"net/http"
	"signin-go/global/logger"

	"github.com/gin-gonic/gin"
)

func HTTPServer() *gin.Engine {
	engine := gin.New()
	engine.GET("/", func(c *gin.Context) {
		logger.Logger.Error(
			"err occurs",
			logger.WrapMeta(
				nil,
				logger.NewMeta("para1", "value1"),
				logger.NewMeta("para2", "value2"),
			)...,
		)
		c.String(http.StatusOK, "Gin Server")
	})
	return engine
}
