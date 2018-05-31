package main

import (
	"log"
	"os"

	"github.com/johananl/gophercon-eu-workshop/pkg/routing"
	"github.com/johananl/gophercon-eu-workshop/webserver"
)

func main() {
	log.Printf("Service is starting...")

	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Port must be specified")
	}

	r := routing.BaseRouter()
	ws := webserver.New("", port, r)

	log.Fatal(ws.Start())
}
