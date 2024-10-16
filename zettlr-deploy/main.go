package main

import (
	"flag"
	"log"
	"os"
)

var (
	InfoFile, _ = os.CreateTemp("", "zettlr-*.info")

	InfoLogger  = log.New(InfoFile, "[Info] ", log.LstdFlags)
	DebugLogger = log.New(os.Stdout, "[Debug] ", log.LstdFlags)
	ErrorLogger = log.New(os.Stdout, "[Error] ", log.LstdFlags)
)

// the expected use case is that the Zettlr export command will call this executable and
// we will presumably be able to use what it gives us to publish a document to our blog.
//
// chettlr-deploy <file_path>

func main() {
	// we begin by checking if the config file contains the information
	// needed to connect to the chettlr database.
	var DatabaseConf DatabaseConfiguration
	err := loadDatabaseConf("/home/diego/.chettlr.json", &DatabaseConf)
	if err != nil {
		ErrorLogger.Fatal(err.Error())
	}
	DebugLogger.Printf("Connecting to %s...\n", DatabaseConf.DatabaseName)

	// connecting to the database
	db, err := getDatabase(DatabaseConf)
	if err != nil {
		ErrorLogger.Fatal(err.Error())
	}
	defer db.Close()

	// test that the database has already been set-up
	// by checking for the maintenance table...
	if !hasMaintenance(db) {
		InfoLogger.Printf("missing maintenance table, going to init the db\n")

		// create the tables the database requires to operate...
		err = createTables(db)
		if err != nil {
			panic(err)
		}
	}

	// collect the html file that was exported and place it into
	// the database
	flag.Parse()
	InfoLogger.Printf("Publishing %s to the website...", flag.Arg(0))

	// upload the html file that was exported to the database
	// by pushing the html content, with the proper meta tags, to the
	// central repository where it can be pulled down by others.
}
