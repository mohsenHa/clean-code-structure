package healthhandler

import (
	"clean-code-structure/service/healthservice"
)

type Handler struct {
	healthService healthservice.Service
}

func New(healthService healthservice.Service) Handler {
	return Handler{
		healthService: healthService,
	}
}
