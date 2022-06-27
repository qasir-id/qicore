package route

import (
	"github.com/labstack/echo/v4"
	ping "gitlab.com/qasir/gateway-fajri/domain/example/handler"
)

type Handler interface {
	Handle(c echo.Context) (err error)
}

var endpoint = map[string]Handler{
    "ping": ping.NewGetPing(),
}
