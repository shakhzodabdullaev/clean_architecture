package main

import (
	controller "clean_arch/BookStore/Controller"
	service "clean_arch/BookStore/Service"

	"github.com/labstack/echo/v4"
)

func main() {
	echoContext := echo.New()
	service := service.NewBookService(nil)
	controller.NewBookController(echoContext, service)
	echoContext.Logger.Info(echoContext.Start(":8090"))
}
