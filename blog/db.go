package blog

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"errors"
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

func getDatabase(dbconf DatabaseConfiguration, psqlURL string) (*sql.DB, error) {
	var psqlInfo string
	if psqlURL == "" {
		psqlInfo = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s",
			dbconf.Host,
			dbconf.Port,
			dbconf.User,
			dbconf.Password,
			dbconf.DatabaseName,
		)
	} else {
		psqlInfo = psqlURL
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}

