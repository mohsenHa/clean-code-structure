package responseerror

import (
	"errors"

	"github.com/labstack/echo/v4"

	"clean-code-structure/pkg/httpmsg"
	"clean-code-structure/validator"
)

func SendErrorResponse(c echo.Context, err error) error {
	msg, code := httpmsg.Error(err)

	var vErr validator.Error
	if errors.As(err, &vErr) {
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  vErr.Fields,
		})
	}

	return echo.NewHTTPError(code, msg)
}
