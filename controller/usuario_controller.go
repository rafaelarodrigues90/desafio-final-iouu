package controller

import (
	"desafio-final-iouu/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
)

var usuarios = model.Usuarios{
	model.Usuario{},
}

// ListarUsuarios - Listar todos os investidores
func ListarUsuarios(e echo.Context) error {
	jsonFile, err := os.Open("./db/database.json")
	if err != nil {
		fmt.Println(err)
	}

	byteJSON, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	objUsuario := usuarios

	err = json.Unmarshal(byteJSON, &objUsuario)
	if err != nil {
		fmt.Println(err)
	}

	// não lê o valor do ID corretamente
	return e.JSON(http.StatusOK, objUsuario[0:])
}

// BuscarUsuario - Listar apenas um investidor pelo ID
func BuscarUsuario(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	index := 0

	for i := range usuarios {
		if usuarios[i].ID == id {
			index = i

			break
		}
	}
	return e.JSON(http.StatusOK, usuarios[index])
}

// AtualizarUsuario - Modificar o dado de um investidor
func AtualizarUsuario(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	index := 0

	for i := range usuarios {
		if usuarios[i].ID == id {
			index = i

			break
		}
	}

	usuarios[index].Nome = e.FormValue("nome")
	usuarios[index].Sobrenome = e.FormValue("sobrenome")
	usuarios[index].Email = e.FormValue("email")

	return e.JSON(http.StatusOK, usuarios[index])
}

// DeletarUsuario - Deletar um investidor
func DeletarUsuario(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	index := 0

	for i := range usuarios {
		if usuarios[i].ID == id {
			index = i

			break
		}
	}

	deletarUsuario := usuarios[index]
	usuarios = append(usuarios[:index], usuarios[index+1:]...)

	return e.JSON(http.StatusOK, deletarUsuario)
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
