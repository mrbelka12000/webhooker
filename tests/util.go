package tests

import (
	"database/sql"
	"log"

	"github.com/mrbelka12000/webhooker/pkg/config"
)

func flushDB(db *sql.DB) {

	for _, query := range queryS {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
}

var queryS = []string{
	`
	DELETE FROM web_hooks;
`,
}

func getDefaultConfig() config.Config {
	return config.Config{
		PGUrl: "postgres://postgres:check@localhost:5432/database?sslmode=disable",
	}
}
