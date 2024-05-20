package healthhandler

import (
	"clean-code-structure/service/healthservice"
	"clean-code-structure/validator/healthvalidator"
)

type Handler struct {
	healthValidator healthvalidator.Validator
	healthService   healthservice.Service
}

func New(healthService healthservice.Service, healthValidator healthvalidator.Validator) Handler {
	return Handler{
		healthService:   healthService,
		healthValidator: healthValidator,
	}
}
