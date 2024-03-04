package service

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"time"

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
	err := wh.validateCU(data)
	if err != nil {
		return err
	}

	return wh.whr.Create(ctx, data)
}

func (wh *WebHooker) List(ctx context.Context, pars models.DataListPars) ([]models.Data, error) {
	return wh.whr.List(ctx, pars)
}

func (wh *WebHooker) validateCU(data *models.Data) error {
	if data.EndTime.Before(time.Now()) {
		return errors.New("bad end time")
	}

	if _, err := url.Parse(data.CallbackURL); err != nil {
		return err
	}

	if data.Method == "" {
		data.Method = http.MethodPost
	}

	return nil
}
