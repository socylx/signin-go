package main

import (
	"net/http"
	"signin-go/router"

	"github.com/jpillora/overseer"
)

func main() {
	overseer.Run(overseer.Config{
		Program: func(state overseer.State) { // prog(state) runs in a child process
			http.Serve(state.Listener, router.HTTPServer())
		},
		Address: ":3000",
		Debug:   true,
	})
}
