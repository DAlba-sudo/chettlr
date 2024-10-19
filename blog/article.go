package blog

import (
	"database/sql"
	"html/template"
)

type Article struct {
	Id          int
	Title       string
	Description string
	Content    	template.HTML 
	Tags        string
}

func ArticleFromRow(row *sql.Rows, article *Article) error {
	err := row.Scan(
		&article.Id,
		&article.Title,
		&article.Description,
		&article.Content,
		&article.Tags,
	)
	return err

}
