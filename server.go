package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DAlba-sudo/chettlr/blog"
)

type WebServer struct {
	address string
	port    int
}

func (ws WebServer) Serve(ctx context.Context) error {
	// generate the address, port pair
	address := fmt.Sprintf("%s:%d", ws.address, ws.port)

	http.Handle("/published/", http.StripPrefix("/published", http.FileServer(http.Dir("./blog/published/"))))
	http.Handle("/blog/", http.StripPrefix("/blog", blog.GetMux()))

	// serve the server
	err := http.ListenAndServe(address, nil)
	return err
}
