package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/johananl/gophercon-eu-workshop/pkg/routing"
	"github.com/johananl/gophercon-eu-workshop/version"
	"github.com/johananl/gophercon-eu-workshop/webserver"
)

func main() {
	// TODO The values don't show - need to troubleshoot.
	log.Printf("Service is starting. version is %s, commit is %s, time is %s",
		version.Release, version.Commit, version.BuildTime,
	)

	shutdown := make(chan error, 2)

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
		err := diagServer.Start()
		shutdown <- err
	}()

	go func() {
		err := ws.Start()
		shutdown <- err
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case killSignal := <-interrupt:
		log.Printf("Got %s. Stopping...", killSignal)
	case err := <-shutdown:
		if err != nil {
			log.Printf("Got an error '%s'. Stopping...", err)
		}
	}

	err := ws.Stop()
	if err != nil {
		log.Println(err)
	}

	err = diagServer.Stop()
	if err != nil {
		log.Println(err)
	}
}
