package tmpl_service

func GetSampleTmpl() string {
	return `package {{.}}service

import (
	"clean-code-structure/param/{{.}}param"
)

// Service layer MUST return  validator.Error or richerror

func (s Service) Sample(req {{.}}param.SampleRequest) ({{.}}param.SampleResponse, error) {

	if err := s.validator.ValidateSampleRequest(req); err != nil {
		return {{.}}param.SampleResponse{}, err
	}

	return {{.}}param.SampleResponse{Message: "everything is good!"}, nil
}
`
}

func GetServiceTmpl() string {
	return `package {{.}}service

import "clean-code-structure/validator/{{.}}validator"

type Service struct {
	validator {{.}}validator.Validator
}

func New(validator {{.}}validator.Validator) Service {
	return Service{
		validator: validator,
	}
}
`
}
