package httpmsg

import (
	"clean-code-structure/logger"
	"clean-code-structure/pkg/errmsg"
	"clean-code-structure/pkg/richerror"
	"errors"
	"net/http"
)

func Error(err error) (message string, code int) {
	re := richerror.RichError{}
	if errors.As(err, &re) {
		msg := re.Message()
		code = mapKindToHTTPStatusCode(re.Kind())

		// we should not expose unexpected error messages
		if code >= http.StatusInternalServerError {
			logger.Logger.Error(msg)
			msg = errmsg.ErrorMsgSomethingWentWrong
		}

		return msg, code
	}

	return err.Error(), http.StatusBadRequest
}

func mapKindToHTTPStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindForbidden:
		return http.StatusForbidden
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
