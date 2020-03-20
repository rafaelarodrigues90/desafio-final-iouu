package controller

import (
	"desafio-final-iouu/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

var usuarios = model.Usuarios{
	model.Usuario{},
}

// ListarUsuarios - Listar todos os investidores
func ListarUsuarios(e echo.Context) error {
	return e.JSON(http.StatusOK, usuarios)
}

// BuscarUsuario - Listar apenas um investidor pelo ID
func BuscarUsuario(e echo.Context) error {
	return e.String(http.StatusOK, "[ok] Buscar usuário")
}

// AtualizarUsuario - Modificar o dado de um investidor
func AtualizarUsuario(e echo.Context) error {
	return e.String(http.StatusOK, "[ok] Atualizar usuário")
}

// DeletarUsuario - Deletar um investidor
func DeletarUsuario(e echo.Context) error {
	return e.String(http.StatusOK, "[ok] Deletar usuário")
}

// CriarUsuario - Adicionar um investidor
func CriarUsuario(e echo.Context) error {
	id, _ := strconv.Atoi(e.FormValue("id"))
	novoUsuario := model.Usuario{
		ID:        id,
		Nome:      e.FormValue("nome"),
		Sobrenome: e.FormValue("sobrenome"),
		Email:     e.FormValue("email"),
	}

	usuarios = append(usuarios, novoUsuario)

	return e.JSON(http.StatusCreated, novoUsuario)
}