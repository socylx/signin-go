package main

import (
	"net/http"
	"signin-go/global"
	"signin-go/global/config"
	"signin-go/global/time"
	"signin-go/router"

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
