package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	ErrNoConfigFile = errors.New("failed to find chettlr config file at ~/.chettlr.json")
)

type DatabaseConfiguration struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
}

func loadDatabaseConf(path string, conf *DatabaseConfiguration) error {
	data, err := os.ReadFile(path)
	if err == os.ErrNotExist {
		return ErrNoConfigFile
	} else if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, conf)
	if err != nil {
		panic(err)
	}

	return nil
}

func getDatabase(dbconf DatabaseConfiguration) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		dbconf.Host,
		dbconf.Port,
		dbconf.User,
		dbconf.Password,
		dbconf.DatabaseName,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}

func hasMaintenance(db *sql.DB) bool {
	rows, err := db.Query("SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = $1;", "maintenance")
	if err != nil {
		panic(err)
	}

	return rows.Next()
}

func createTables(db *sql.DB) error {
	
	return nil
}