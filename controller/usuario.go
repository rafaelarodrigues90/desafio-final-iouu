package controller

import (
	"net/http"
	"github.com/labstack/echo"
)

func teste (e echo.Context) error {
	return e.String(http.StatusOK, "Hello world")
}

// Listar todos os investidores
func listarUsuarios(e echo.Context) error {
	return e.String(http.StatusOK, "[ok] Listar usuários")
}

// Listar apenas um investidor pelo ID
func buscarUsuario(e echo.Context) error {
	return e.String(http.StatusOK, "[ok] Buscar usuário")
}

// Modificar o dado de um investidor
func atualizarUsuario(e echo.Context) error {
	return e.String(http.StatusOK, "[ok] Atualizar usuário")
}

// Deletar um investidor
func deletarUsuario(e echo.Context) error {
	return e.String(http.StatusOK, "[ok] Deletar usuário")
}

// Adicionar um investidor
func criarUsuario(e echo.Context) error {
	return e.String(http.StatusOK, "[ok] Criar usuário")
}