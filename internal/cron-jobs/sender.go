package cron_jobs

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/rs/zerolog"

	"github.com/mrbelka12000/webhooker/internal/models"
	"github.com/mrbelka12000/webhooker/internal/service"
)

const (
	defaultLimit   = 100
	defaultTimeout = 30 * time.Second
)

type webhooker struct {
	client *http.Client
	srv    *service.WebHooker
	log    zerolog.Logger
}

func newWH(srv *service.WebHooker, log zerolog.Logger) *webhooker {
	return &webhooker{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
		srv: srv,
		log: log,
	}
}

func (wh *webhooker) send() {
	hooks, err := wh.srv.List(context.Background(), models.DataListPars{
		Limit: defaultLimit,
	})
	if err != nil {
		return
	}

	for _, hook := range hooks {
		if hook.EndTime.Before(time.Now()) {
			parsedURL, err := url.Parse(hook.CallbackURL)
			if err != nil {
				wh.log.Err(err).Msg("parse url")
				return
			}

			params, _ := url.ParseQuery(parsedURL.RawQuery)
			for k, v := range hook.Params {
				params.Add(k, v[0])
			}

			parsedURL.RawQuery = params.Encode()

			req, err := http.NewRequest(hook.Method, parsedURL.String(), strings.NewReader(hook.Body))
			if err != nil {
				wh.log.Err(err).Msg("create request")
				continue
			}

			_, err = wh.client.Do(req)
			if err != nil {
				wh.log.Err(err).Msg("send request")
				continue
			}

			wh.log.Info().Msg(fmt.Sprintf("notifications sended to url %s", parsedURL.String()))
		}
	}
}
