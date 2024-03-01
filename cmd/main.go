package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/rs/zerolog"

	"github.com/mrbelka12000/webhooker/internal/handler"
	"github.com/mrbelka12000/webhooker/internal/repo"
	"github.com/mrbelka12000/webhooker/internal/service"
	"github.com/mrbelka12000/webhooker/pkg/config"
	"github.com/mrbelka12000/webhooker/pkg/database"
	"github.com/mrbelka12000/webhooker/pkg/server"
)

func main() {

	log := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	cfg, err := config.Get()
	if err != nil {
		log.Err(err).Msg("get config")
		return
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Err(err).Msg("connect to database")
		return
	}

	repo := repo.NewWebHooker(db, log)
	srv := service.NewWebHooker(repo, log)
	router := handler.RegisterRoutes(srv)
	httpServer := server.NewServer(router, cfg)

	log.Info().Msg(fmt.Sprintf("server started on port %s", cfg.HttpPort))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	select {
	case s := <-done:
		log.Info().Msg("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Err(err).Msg("http server notify")
	}
}
