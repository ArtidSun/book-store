package main

import (
	"book-store/connections"
	"book-store/route"
	book_handler "book-store/services/http"
	book_repo "book-store/services/repository"
	book_usecase "book-store/services/usecase"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	conn, err := connections.InitializeDatabase()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	bookRepo := book_repo.NewRepository(conn)
	bookUs := book_usecase.NewBookUsecase(bookRepo)
	bookHandler := book_handler.NewBookHandler(bookUs)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	router := route.NewRoute(e)
	router.NewBookRoute(bookHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
