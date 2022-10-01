package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jpillora/overseer"
)

// prog(state) runs in a child process
func prog(state overseer.State) {
	log.Printf("app (%s) listening...", state.ID)
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "app (%s) says hello\n", state.ID)
	}))
	http.Serve(state.Listener, nil)
}

// create another main() to run the overseer process
// and then convert your old main() into a 'prog(state)'
func main() {
	overseer.Run(overseer.Config{
		Program: prog,
		Address: ":3000",
	})
}
