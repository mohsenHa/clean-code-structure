package healthparam

import (
	"clean-code-structure/param"
)

type CheckRequest struct {
	param.BaseRequest
}

type CheckResponse struct {
	param.BaseResponse
	Message string `json:"message"`
}
