package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	// iniciando echo
	e := echo.New()

	// função de teste
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "Teste ok")
	})

	// iniciando servidor na porta 8000
	e.Logger.Fatal(e.Start(":8000"))
}
