package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

// Index ...
func Index() *echo.Echo {
	// iniciando echo
	e := echo.New()

	// teste
	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "Funfou!")
	})

	// Grupo de rotas
	RotaUsuario(e.Group("/usuarios"))

	return e
}
