package main

import (
	"desafio-final-iouu/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// iniciando echo
	e := routes.Index()

	// middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// iniciando servidor na porta 8000
	e.Logger.Fatal(e.Start(":8000"))
}
