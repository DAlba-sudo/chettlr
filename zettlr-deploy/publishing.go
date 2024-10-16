package main

import (
	"database/sql"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"
)

type YAMLFrontmatter struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

// for a given Zettlr markdown file, find the corresponding HTML file that we wish to upload the contents
// of to the database.
func findTargetHTML(target_path string) string {
	dir, file := path.Split(target_path)
	file_name := strings.Split(file, ".")[0]

	entries, err :=  os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		entry_name := strings.Split(entry.Name(), ".")[0]
		if strings.Compare(file_name, entry_name) == 0 {
			return path.Join(dir, entry.Name())
		}
	}

	return ""
}

// parses the YAML frontmatter for the title, description, and content.
func parseYamlFrontmatter(target_path string, frontmatter *YAMLFrontmatter) error {
	data, err := os.ReadFile(target_path)
	if err != nil {
		panic(err)
	}

	is_yaml := false
	yaml_content := &strings.Builder{}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.Contains(line, "---") {
			is_yaml = !is_yaml
		} else if is_yaml {
			yaml_content.WriteString(line + "\n")
		} 
	}

	// we now parse the provided yaml content for the relevant items
	// to parse.
	err = yaml.Unmarshal([]byte(yaml_content.String()), frontmatter)	
	if err != nil {
		panic(err)
	}

	return nil
}

func publish(db *sql.DB, html_path string, frontmatter *YAMLFrontmatter) error {
	html, err := os.ReadFile(html_path)
	if err != nil {
		return err
	}	

	_, err = db.Exec("INSERT INTO articles (title, description, content) VALUES ($1, $2, $3);", 
		frontmatter.Title,
		frontmatter.Description,
		html,
	)
	if err != nil {
		return err
	}

	return nil
}