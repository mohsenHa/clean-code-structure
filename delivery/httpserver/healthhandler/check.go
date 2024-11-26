package healthhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"clean-code-structure/param/healthparam"
	"clean-code-structure/pkg/responseerror"
)

func (h Handler) healthCheck(c echo.Context) error {
	var req healthparam.CheckRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.healthService.Check(req)
	if err != nil {
		return responseerror.SendErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}
