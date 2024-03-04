package tests

import (
	"github.com/rs/zerolog"

	"github.com/mrbelka12000/webhooker/internal/service"
)

var app struct {
	uc  *service.WebHooker
	log zerolog.Logger
}
