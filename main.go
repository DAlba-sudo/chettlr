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
	port := flag.Int("port", 8443, "port to use when binding to the server")
	flag.Parse()
	main_supervisor := suture.NewSimple("WebServer")
	
	ws := WebServer{
		address: "0.0.0.0",
		port:    *port,
	}
	_ = main_supervisor.Add(ws)

	fmt.Println("Starting webserver...")
	err := main_supervisor.Serve(ctx)
	if err != nil {
		panic(err)
	}
}
