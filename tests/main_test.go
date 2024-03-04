// Package tests implements integration tests.
package tests

import (
	"os"
	"testing"

	"github.com/rs/zerolog"

	"github.com/mrbelka12000/webhooker/internal/repo"
	"github.com/mrbelka12000/webhooker/internal/service"
	"github.com/mrbelka12000/webhooker/pkg/database"
)

func TestMain(m *testing.M) {
	log := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	cfg := getDefaultConfig()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Err(err).Msg("connect to database")
		return
	}
	defer db.Close()

	flushDB(db)

	whr := repo.NewWebHooker(db, log)
	app.uc = service.NewWebHooker(whr, log)

	os.Exit(m.Run())
}
