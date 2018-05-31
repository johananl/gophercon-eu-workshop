package main

import (
	"log"
	"os"

	"github.com/johananl/gophercon-eu-workshop/pkg/routing"
	"github.com/johananl/gophercon-eu-workshop/version"
	"github.com/johananl/gophercon-eu-workshop/webserver"
)

func main() {
	log.Printf("Service is starting. version is %s, commit is %s, time is %s",
		// TODO The values don't show - need to troubleshoot.
		version.Release, version.Commit, version.BuildTime,
	)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("PORT must be specified")
	}

	r := routing.BaseRouter()
	ws := webserver.New("", port, r)

	internalPort := os.Getenv("INTERNAL_PORT")
	if len(internalPort) == 0 {
		log.Fatal("INTERNAL_PORT must be specified")
	}
	diagRouter := routing.DiagnosticsRouter()
	diagServer := webserver.New("", internalPort, diagRouter)

	go func() {
		log.Fatal(diagServer.Start())
	}()

	log.Fatal(ws.Start())
}
