package cron_jobs

import (
	"context"
	"net/http"
	"time"

	"github.com/mrbelka12000/webhooker/internal/models"
	"github.com/mrbelka12000/webhooker/internal/service"
)

const defaultLimit = 100

type webhooker struct {
	client *http.Client
	srv    *service.WebHooker
}

func (wh *webhooker) sender() {
	hooks, err := wh.srv.List(context.Background(), models.DataListPars{
		Limit: 100,
	})
	if err != nil {
		return
	}

	for _, hook := range hooks {
		if hook.EndTime.Before(time.Now()) {

		}
	}
}
