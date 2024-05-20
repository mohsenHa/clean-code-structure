package param

import (
	"context"
	"net/http"
)

type BaseRequest struct {
	Ctx      context.Context
	Request  *http.Request
	Response *http.Response
}
type BaseResponse struct {
}
