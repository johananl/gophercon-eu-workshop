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
		log.Fatal("Port must be specified")
	}

	r := routing.BaseRouter()
	ws := webserver.New("", port, r)

	log.Fatal(ws.Start())
}
