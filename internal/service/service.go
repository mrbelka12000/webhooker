package service

import (
	"github.com/rs/zerolog"

	"github.com/mrbelka12000/webhooker/internal/repo"
)

type WebHooker struct {
	whr *repo.WebHooker
	log zerolog.Logger
}

func NewWebHooker(whr *repo.WebHooker, log zerolog.Logger) *WebHooker {
	return &WebHooker{
		whr: whr,
		log: log,
	}
}
