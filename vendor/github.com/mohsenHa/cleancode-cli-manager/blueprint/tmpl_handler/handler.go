package tmpl_handler

func GetHandlerTmpl() string {
	return `package {{.}}handler

import (
	"clean-code-structure/service/{{.}}service"
)

type Handler struct {
	{{.}}Service   {{.}}service.Service
}

func New({{.}}Service {{.}}service.Service) Handler {
	return Handler{
		{{.}}Service:   {{.}}Service,
	}
}
`
}
func GetRoutTmpl() string {
	return `package {{.}}handler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(messageGroup *echo.Group) {

	messageGroup.GET("/sample", h.sample)

}
`
}
func GetSampleTmpl() string {
	return `package {{.}}handler

import (
	"clean-code-structure/param/healthparam"
	"clean-code-structure/pkg/responseerror"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) sample(c echo.Context) error {
	var req {{.}}param.SampleRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.{{.}}Service.Sample(req)

	if err != nil {
		return responseerror.SendErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}
`
}
