//Go algorithm to add static files in directory to a server
//Implementation 2 of GoServe
//Adds all static file pings requested at runtime as handled by other http requests
//

package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))

	http.HandleFunc("/ping", InitialPing)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func InitialPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TEST"))
}
