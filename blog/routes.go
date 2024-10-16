package blog

import (
	"database/sql"
	"html/template"
	"net/http"
)

var (
	Mux = http.NewServeMux()
	DatabaseConf DatabaseConfiguration
	db *sql.DB
)

func GetMux() http.Handler {
	loadDatabaseConf("/home/diego/.chettlr.json", &DatabaseConf)
	ldb, err := getDatabase(DatabaseConf)	
	if err != nil {
		panic(err)
	}
	db = ldb

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

	titles := []string{}
	rows, err := db.Query("SELECT * FROM articles LIMIT 10;")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var title string
		var description string
		var content string
		var tags string

		rows.Scan(
			&id,
			&title,
			&description,
			&content,
			&tags,
		)

		titles = append(titles, title)
	}


	templ.Execute(wr, titles)
}