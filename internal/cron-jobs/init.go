package cron_jobs

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog"

	"github.com/mrbelka12000/webhooker/internal/service"
)

func Start(srv *service.WebHooker, log zerolog.Logger) {
	s := gocron.NewScheduler(time.UTC)
	wh := newWH(srv, log)

	s.Every(10 * time.Second).Do(func() {
		wh.send()
	})

	s.StartBlocking()
}
