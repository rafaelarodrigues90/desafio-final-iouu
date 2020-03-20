package routes

import (
	"desafio-final-iouu/controller"

	"github.com/labstack/echo"
)

// RotaUsuario ...
func RotaUsuario(g *echo.Group) {

	// Grupo de rotas - teste
	g.GET("/usuarios", controller.ListarUsuarios)
	g.GET("/usuarios/:id", controller.ListarUsuarios)
	g.PUT("/atualizar-usuario/:id", controller.AtualizarUsuario)
	g.DELETE("/deletar-usuario/:id", controller.DeletarUsuario)
	g.POST("/criar-usuario", controller.CriarUsuario)

}