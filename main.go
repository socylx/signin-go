package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
)

// prog(state) runs in a child process
func prog(state overseer.State) {
	engine := gin.New()
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Gin Server")
	})
	http.Serve(state.Listener, engine)
}

// create another main() to run the overseer process
// and then convert your old main() into a 'prog(state)'
func main() {
	overseer.Run(overseer.Config{
		Program: prog,
		Address: ":3000",
		Debug:   true,
	})
}
