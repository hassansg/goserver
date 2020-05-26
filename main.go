//Go algorithm for displaying strings to servers
//Implementation 1 of GoServe
//Handles a http request with a string in the header to display to the user.
//

package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", Initial)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

	http.HandleFunc("/", UsrInput)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func Initial(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Relay " + message
	w.Write([]byte(message))
}

func UsrInput(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	fmt.Println("Enter string to be displayed on the server: ")
	var i1 string
	fmt.Scanf("%s", &i1)
	w.Write([]byte(i1))
}
