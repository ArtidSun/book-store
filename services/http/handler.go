package http

import (
	"book-store/services"
	"net/http"
	"strconv"
	"sync"

	helperModel "git.innovasive.co.th/backend/models"
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

func (b *bookHandler) FetchListBooks(c echo.Context) error {
	ctx := c.Request().Context()
	title := c.QueryParam("search_word")
	minPrice := c.QueryParam("min_price")
	maxPrice := c.QueryParam("max_price")
	var page, pageErr = strconv.Atoi(c.QueryParam("page"))
	var perPage, perPageErr = strconv.Atoi(c.QueryParam("per_page"))
	var paginator = helperModel.NewPaginator()

	if pageErr == nil {
		paginator.Page = page
	}

	if perPageErr == nil {
		paginator.PerPage = perPage
	}
	var args = new(sync.Map)

	if title != "" {
		args.Store("search_word", title)
	}
	if minPrice != "" {
		args.Store("min_price", title)
	}
	if maxPrice != "" {
		args.Store("max_price", title)
	}

	books, err := b.bookUs.FetchListBooks(ctx, args, &paginator)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := map[string]interface{}{
		"books": books,
	}

	return c.JSON(http.StatusOK, response)
}
