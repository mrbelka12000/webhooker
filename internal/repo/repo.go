package repo

import (
	"database/sql"

	"github.com/rs/zerolog"
)

type WebHooker struct {
	db *sql.DB
	lg zerolog.Logger
}

func NewWebHooker(db *sql.DB, log zerolog.Logger) *WebHooker {
	return &WebHooker{
		db: db,
		lg: log,
	}
}
