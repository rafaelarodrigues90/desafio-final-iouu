package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Index ...
func Index() *echo.Echo {
	// iniciando echo
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())

	// teste
	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "Funfou!")
	})

	// Grupo de rotas
	RotaUsuario(e.Group("/usuarios"))

	return e
}
