package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/thejerf/suture/v4"
)

var (
	ctx = context.Background()
)

func main() {
	flag.Parse()
	main_supervisor := suture.NewSimple("WebServer")
	
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		panic(err)
	}

	// add the webserver as a service
	ws := WebServer{
		address: "0.0.0.0",
		port:    port,
	}
	_ = main_supervisor.Add(ws)

	fmt.Println("Starting webserver...")
	err = main_supervisor.Serve(ctx)
	if err != nil {
		panic(err)
	}
}
