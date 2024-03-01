package service

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/mrbelka12000/webhooker/internal/models"
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

func (wh *WebHooker) Create(ctx context.Context, data *models.Data) error {
	return wh.whr.Create(ctx, data)
}

func (wh *WebHooker) List(ctx context.Context, pars models.DataListPars) ([]models.Data, error) {
	return wh.whr.List(ctx, pars)
}
