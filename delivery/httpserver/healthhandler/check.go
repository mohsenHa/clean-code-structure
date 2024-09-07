package healthhandler

import (
	"clean-code-structure/param/healthparam"
	"clean-code-structure/pkg/httpmsg"
	"clean-code-structure/validator"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) healthCheck(c echo.Context) error {
	var req healthparam.CheckRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.healthService.Check(req)
	var vErr validator.Error

	if errors.As(err, &vErr) {
		msg, code := httpmsg.Error(err)

		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  vErr.Fields,
		})
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
