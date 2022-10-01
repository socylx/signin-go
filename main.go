package main

import (
	"log"
	"net/http"
	"signin-go/global"
	"signin-go/global/config"
	"signin-go/router"

	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
)

func main() {
	global.PreInit()

	overseer.Run(overseer.Config{
		Program: func(state overseer.State) { // prog(state) runs in a child process
			defer func() {
				config.Server = nil
				log.Println("config.Server: ", config.Server)
				// TODO: 关闭一些东西
			}()
			global.Init()
			http.Serve(state.Listener, router.HTTPServer())
		},
		Address: config.Server.Port,
		Debug:   config.Server.Mode != gin.ReleaseMode,
	})

}
