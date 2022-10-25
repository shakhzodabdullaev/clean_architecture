package main

import (
	controller "clean_arch/BookStore/Controller"
	service "clean_arch/BookStore/Service"
	"clean_arch/BookStore/database"

	"github.com/labstack/echo/v4"
)

func main() {
	echoContext := echo.New()
	datalayer := database.NewBookDatalayerImpl()
	service := service.NewBookServiceImpl(datalayer)
	controller.NewBookController(echoContext, service)
	echoContext.Logger.Info(echoContext.Start(":8090"))
}
