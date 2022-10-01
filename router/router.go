package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HTTPServer() *gin.Engine {
	engine := gin.New()
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Gin Server")
	})
	return engine
}
