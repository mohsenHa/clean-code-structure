package healthhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(messageGroup *echo.Group) {

	messageGroup.GET("/check", h.healthCheck)

}
