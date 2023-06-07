package services

import "github.com/labstack/echo/v4"

type HandlerInterface interface {
	HelloWorld(c echo.Context) error
	FetchListBooks(c echo.Context) error
}
