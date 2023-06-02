package http

import (
	"book-store/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bookHandler struct {
	bookUs services.UsecaseInterface
}

func NewBookHandler(bookUs services.UsecaseInterface) services.HandlerInterface {
	return &bookHandler{
		bookUs: bookUs,
	}
}

func (b *bookHandler) HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello-World")
}