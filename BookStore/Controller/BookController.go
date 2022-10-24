package controller

import (
	"clean_arch/BookStore/intrface"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	bookService intrface.BookService
}

func NewBookController(echoCtx *echo.Echo, bookServiceObject intrface.BookService) {
	bookControllerObject := &BookController{
		bookService: bookServiceObject,
	}

	echoCtx.GET("/printAuthor", bookControllerObject.printAuthor)
}

func (book *BookController) printAuthor(ec echo.Context) {

}
