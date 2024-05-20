package healthservice

import (
	"clean-code-structure/param/healthparam"
)

func (s Service) Check(req healthparam.CheckRequest) (healthparam.CheckResponse, error) {
	return healthparam.CheckResponse{Message: "everything is good!"}, nil
}
