package main

import (
	"flag"
	"io"
	"log"
	"os"
)

var (
	logfile, _ = os.CreateTemp("", "zettlr-*.log");
	logOut = io.MultiWriter(os.Stdout, logfile)

	InfoLogger  = log.New(logOut, "[Info] ", log.LstdFlags)
	DebugLogger = log.New(logOut, "[Debug] ", log.LstdFlags)
	ErrorLogger = log.New(logOut, "[Error] ", log.LstdFlags)
)

// the expected use case is that the Zettlr export command will call this executable and
// we will presumably be able to use what it gives us to publish a document to our blog.
//
// chettlr-deploy <file_path>

func main() {
	flag.Parse()

	var DatabaseConf DatabaseConfiguration

	if (flag.Arg(0) == "") {
		// we begin by checking if the config file contains the information
		// needed to connect to the chettlr database.
		err := loadDatabaseConf("/home/diego/.chettlr.json", &DatabaseConf)
		if err != nil {
			ErrorLogger.Fatal(err.Error())
		}
	} else {
		
	}

	DebugLogger.Printf("Connecting to %s...\n", DatabaseConf.DatabaseName)

	// connecting to the database
	db, err := getDatabase(DatabaseConf, flag.Arg(0))
	if err != nil {
		ErrorLogger.Fatal(err.Error())
	}
	defer db.Close()

	// test that the database has already been set-up
	// by checking for the maintenance table...
	if !hasTable(db, "articles") {
		InfoLogger.Printf("missing article table, going to init the db\n")

		// create the tables the database requires to operate...
		err = createTables(db)
		if err != nil {
			panic(err)
		}
	}

	// collect the html file that was exported and place it into
	// the database
	flag.Parse()

	publish_target := flag.Arg(1)
	InfoLogger.Printf("Publishing %s to the website...", publish_target)

	// upload the html file that was exported to the database
	// by pushing the html content, with the proper meta tags, to the
	// central repository where it can be pulled down by others.
	html_path := findTargetHTML(publish_target)
	if html_path != "" {
		DebugLogger.Printf("found the publishing target HTML file (%s)", html_path)
	} else {
		ErrorLogger.Fatalf("could not find the target html file for %s", publish_target)
	}

	var frontmatter YAMLFrontmatter
	err = parseYamlFrontmatter(publish_target, &frontmatter)
	if err != nil {
		panic(err)
	}

	DebugLogger.Printf("parsed yaml (title: %s, desc: %s)", frontmatter.Title, frontmatter.Description)

	err = publish(db, html_path, &frontmatter)
	if err != nil {
		panic(err)
	}

	InfoLogger.Print("article has been published")
}
