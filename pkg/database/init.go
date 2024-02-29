package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/mrbelka12000/webhooker/pkg/config"
)

// Connect ..
func Connect(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.PGUrl)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return db, nil
}
