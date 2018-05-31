package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Printf("Service is starting...")
	http.HandleFunc("/home", homeHandler())
	http.ListenAndServe(":8080", nil)
}

func homeHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request is processing: %s", r.URL.Path)
		fmt.Fprint(w, "Hello! Your request was processed.")
	}
}
