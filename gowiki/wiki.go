package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %v\n", r.URL.Path)
		log.Println(r.URL.Path)
	})

	http.HandleFunc("/wiki", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: wiki")
	})

	http.ListenAndServe(":8080", nil)
}
