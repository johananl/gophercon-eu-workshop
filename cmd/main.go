package main

import (
	"log"
	"net/http"

	"github.com/johananl/gophercon-eu-workshop/pkg/routing"
)

func main() {
	log.Printf("Service is starting...")

	r := routing.BaseRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}
