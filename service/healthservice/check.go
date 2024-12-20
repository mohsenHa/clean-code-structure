package healthservice

import (
	"clean-code-structure/param/healthparam"
)

// Service layer MUST return  validator.Error or richerror

func (s Service) Check(req healthparam.CheckRequest) (healthparam.CheckResponse, error) {
	if err := s.validator.ValidateCheckRequest(req); err != nil {
		return healthparam.CheckResponse{}, err
	}

	return healthparam.CheckResponse{Message: "everything is good!"}, nil
}
