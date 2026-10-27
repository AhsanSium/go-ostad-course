package main

import (
	"fmt"
	"net/http"
)

func main () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, From Go, You Requested: %s\n", r.URL.Path)
	})
	fmt.Println("Starting Server at 8010")
	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
