package main

import (
	"log"
	"net/http"
	"signin-go/global"
	"signin-go/global/config"
	"signin-go/router"
	"strconv"

	"github.com/jpillora/overseer"
)

func main() {
	global.Init()

	log.Println("config.Server1: ", config.Server)

	overseer.Run(overseer.Config{
		Program: func(state overseer.State) { // prog(state) runs in a child process
			defer func() {
				config.Server = nil
				log.Println("config.Server: ", config.Server)
				// TODO: 关闭一些东西
			}()
			log.Println("Program: ", config.Server)
			http.Serve(state.Listener, router.HTTPServer())
		},
		Address: ":" + strconv.Itoa(config.Server.Port),
		Debug:   true,
	})

	log.Println("config.Server2: ", config.Server)
}
