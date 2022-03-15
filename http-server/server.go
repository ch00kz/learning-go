package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got request")
		fmt.Println(r)
		fmt.Fprintf(w, "Hello world", r.URL.Path)
	})

	fmt.Print("Starting server on port 3000")

	http.ListenAndServe(":3000", nil)
}
