package healthservice

import "clean-code-structure/validator/healthvalidator"

type Service struct {
	validator healthvalidator.Validator
}

func New(validator healthvalidator.Validator) Service {
	return Service{
		validator: validator,
	}
}
