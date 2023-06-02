package route

import (
	"book-store/services"

	"github.com/labstack/echo/v4"
)

type Route struct {
	e *echo.Echo
}

func NewRoute(e *echo.Echo) *Route {
	return &Route{e: e}
}

func (r *Route) NewBookRoute(handler services.HandlerInterface) {
	r.e.GET("/hello-world", handler.HelloWorld)
}
