package cron_jobs

import (
	"time"

	"github.com/go-co-op/gocron"

	"github.com/mrbelka12000/webhooker/internal/service"
)

func Start(srv *service.WebHooker) {
	s := gocron.NewScheduler(time.UTC)

	s.StartBlocking()
}
