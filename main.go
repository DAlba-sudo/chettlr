package main

import (
	"context"
	"fmt"

	"github.com/thejerf/suture/v4"
)

var (
	ctx = context.Background()
)

func main() {

	main_supervisor := suture.NewSimple("WebServer")

	// add the webserver as a service
	ws := WebServer{
		address: "127.0.0.1",
		port:    8443,
	}
	_ = main_supervisor.Add(ws)

	fmt.Println("Starting webserver...")
	err := main_supervisor.Serve(ctx)
	if err != nil {
		panic(err)
	}
}
