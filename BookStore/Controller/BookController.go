package controller

import (
	"clean_arch/BookStore/intrface"
	"fmt"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	bookService intrface.BookService
}

func NewBookController(echo *echo.Echo, bookServiceObject intrface.BookService) {
	bookControllerObject := &BookController{
		bookService: bookServiceObject,
	}

	echo.GET("/printAuthor", bookControllerObject.PrintAuthor)
	echo.GET("/test", bookControllerObject.Test)
}

func (controllerObj *BookController) PrintAuthor(ec echo.Context) error {

	return nil
}

func (controllerObj *BookController) Test(ec echo.Context) error {
	fmt.Println("****** Inside Book Controller ******")
	requestContext := ec.Request().Context()
	controllerObj.bookService.TestBookService(requestContext)
	return nil
}
