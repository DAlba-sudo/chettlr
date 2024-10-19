package blog

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
)

var (
	Mux          = http.NewServeMux()
	DatabaseConf DatabaseConfiguration
	db           *sql.DB
)

func GetMux() http.Handler {
	if flag.Arg(0) == "" {
		loadDatabaseConf("/home/diego/.chettlr.json", &DatabaseConf)
	}
	ldb, err := getDatabase(DatabaseConf, flag.Arg(0))
	if err != nil {
		panic(err)
	}
	db = ldb

	// initialize all the relevant routes that pertain to the blog part of the
	// website here.
	Mux.HandleFunc("/", handleIndex)
	Mux.HandleFunc("/article/{id}/", handleIndividualArticle)
	Mux.HandleFunc("/article/{id}/content", serveIndividualArticleContent)

	return Mux
}

func serveIndividualArticleContent(wr http.ResponseWriter, req *http.Request) {
	row, err := db.Query("SELECT * FROM articles WHERE id=$1 LIMIT 1;", req.PathValue("id"))
	if row.Next() && err != nil {
		panic(err)
	}
	defer row.Close()

	var article Article
	err = ArticleFromRow(row, &article)
	if err != nil {
		panic(err)
	}

	wr.Write([]byte(article.Content))

}

func handleIndividualArticle(wr http.ResponseWriter, req *http.Request) {
	templ, err := template.ParseFiles("./web/templates/basic.html", "./web/templates/navbar.html", "./blog/templates/individual_article.html")
	if err != nil {
		panic(err)
	}

	row, err := db.Query("SELECT * FROM articles WHERE id=$1 LIMIT 1;", req.PathValue("id"))
	if row.Next() && err != nil {
		panic(err)
	}
	defer row.Close()

	var article Article
	err = ArticleFromRow(row, &article)
	if err != nil {
		panic(err)
	}

	log.Printf("Showing article #%s, title = %s", req.PathValue("id"), article.Title)
	templ.Execute(wr, article)
}

func handleIndex(wr http.ResponseWriter, req *http.Request) {
	templ, err := template.ParseFiles("./web/templates/basic.html", "./web/templates/navbar.html", "./blog/templates/index.html")
	if err != nil {
		panic(err)
	}

	articles := []*Article{}
	rows, err := db.Query("SELECT * FROM articles LIMIT 10;")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		article := &Article{}
		err = ArticleFromRow(rows, article)
		if err != nil {
			panic(err)
		}

		articles = append(articles, article)
	}
	defer rows.Close()

	templ.Execute(wr, articles)
}
