package main

import (
	"gsteps-go/global"
	"gsteps-go/global/config"
	"gsteps-go/global/time"
	"gsteps-go/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
)

func main() {
	config.PreInit()

	overseer.Run(overseer.Config{
		Program: func(state overseer.State) { // prog(state) runs in a child process
			defer func() {
				global.Close()
			}()
			global.Init()
			http.Serve(state.Listener, router.HTTPServer())
		},
		Address:          config.Server.Port,
		TerminateTimeout: 10 * time.Second,
		Debug:            config.Server.Mode != gin.ReleaseMode,
	})

}
