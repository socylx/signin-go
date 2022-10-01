package main

import (
	"log"
	"net/http"
	"signin-go/global"
	"signin-go/global/config"
	"signin-go/global/logger"
	"signin-go/router"

	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
)

func main() {
	config.PreInit()

	overseer.Run(overseer.Config{
		Program: func(state overseer.State) { // prog(state) runs in a child process
			defer func() {
				_ = logger.Logger.Sync()
				log.Println("关闭一些东西2: ")
			}()
			global.Init()
			http.Serve(state.Listener, router.HTTPServer())
		},
		Address: config.Server.Port,
		Debug:   config.Server.Mode != gin.ReleaseMode,
	})

}
