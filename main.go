package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func serveRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Helo world")
}

func main() {
	fmt.Printf("Server listening on port: %s\n", port)

	http.HandleFunc("/", serveRequest)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
