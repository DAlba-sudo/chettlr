package blog

import (
	"html/template"
	"net/http"
)

var (
	Mux = http.NewServeMux()
)

func GetMux() http.Handler {
	// initialize all the relevant routes that pertain to the blog part of the
	// website here.
	Mux.HandleFunc("/", handleIndex)

	return Mux
}

func handleIndex(wr http.ResponseWriter, req *http.Request) {
	templ, err := template.ParseFiles("./web/templates/basic.html", "./blog/templates/index.html")
	if err != nil {
		panic(err)
	}

	templ.Execute(wr, nil)
}