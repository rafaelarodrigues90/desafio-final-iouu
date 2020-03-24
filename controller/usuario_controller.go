package controller

import (
	"desafio-final-iouu/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

var usuarios = model.Usuarios{
	model.Usuario{},
}

// ListarUsuarios - Listar todos os investidores
func ListarUsuarios(e echo.Context) error {
	// lê o json
	jsonFile, err := os.Open("./db/database.json")
	if err != nil {
		fmt.Println("OS ERROR: ", err)
	}

	// armazena os dados em formato byte
	byteJSON, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("IOUTIL ERROR: ", err)
	}

	// recebe modelo
	objUsuario := usuarios

	// transforma byte em string
	err = json.Unmarshal(byteJSON, &objUsuario)
	if err != nil {
		fmt.Println("UNMARSHALL ERROR: ", err)
	}

	return e.JSON(http.StatusOK, objUsuario)
}

// BuscarUsuario - Listar apenas um investidor pelo ID
func BuscarUsuario(e echo.Context) error {
	jsonFile, err := os.Open("./db/database.json")
	if err != nil {
		fmt.Println("OS ERROR: ", err)
	}

	// armazena os dados em formato byte
	byteJSON, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("IOUTIL ERROR: ", err)
	}

	// recebe modelo
	objUsuario := usuarios

	// transforma byte em string
	err = json.Unmarshal(byteJSON, &objUsuario)
	if err != nil {
		fmt.Println("UNMARSHALL ERROR: ", err)
	}

	id := e.Param("id") // pega o valor da string

	index := 0
	for i := range objUsuario {
		// comparando as strings
		if objUsuario[i].ID == id {
			index = i
			return e.JSON(http.StatusOK, objUsuario[index])
		}
	}
	return e.JSON(http.StatusNotFound, "Usuário não existe")
}

// AtualizarUsuario - Modificar o dado de um investidor
func AtualizarUsuario(e echo.Context) error {
	jsonFile, err := os.Open("./db/database.json")
	if err != nil {
		fmt.Println("OS ERROR: ", err)
	}

	// armazena os dados em formato byte
	byteJSON, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("IOUTIL ERROR: ", err)
	}

	// recebe modelo
	objUsuario := usuarios

	// transforma byte em string
	err = json.Unmarshal(byteJSON, &objUsuario)
	if err != nil {
		fmt.Println("UNMARSHALL ERROR: ", err)
	}

	id := e.Param("id") // pega o valor da string

	index := 0
	for i := range objUsuario {
		// comparando as strings
		if objUsuario[i].ID == id {
			index = i
			break
		}
	}

	objUsuario[index].Nome = e.FormValue("nome")
	objUsuario[index].Sobrenome = e.FormValue("sobrenome")
	objUsuario[index].Email = e.FormValue("email")

	byteJSON, err = json.Marshal(objUsuario)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("./db/database.json", byteJSON, 0644)
	if err != nil {
		fmt.Println("WRITEFILE ERROR", err)
	}

	return e.JSON(http.StatusOK, objUsuario[index])
}

// DeletarUsuario - Deletar um investidor
func DeletarUsuario(e echo.Context) error {
	jsonFile, err := os.Open("./db/database.json")
	if err != nil {
		fmt.Println("OS ERROR: ", err)
	}

	// armazena os dados em formato byte
	byteJSON, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("IOUTIL ERROR: ", err)
	}

	// recebe modelo
	objUsuario := usuarios

	// transforma byte em string
	err = json.Unmarshal(byteJSON, &objUsuario)
	if err != nil {
		fmt.Println("UNMARSHALL ERROR: ", err)
	}

	id := e.Param("id")
	index := 0

	for i := range objUsuario {
		if objUsuario[i].ID == id {
			index = i

			break
		}
	}

	deletarUsuario := objUsuario[index]
	usuarios = append(objUsuario[:index], objUsuario[index+1:]...)

	byteJSON, err = json.Marshal(objUsuario)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("./db/database.json", byteJSON, 0644)
	if err != nil {
		fmt.Println("WRITEFILE ERROR", err)
	} else {
		fmt.Println("Usuario deletado", objUsuario[index])
	}

	return e.JSON(http.StatusOK, deletarUsuario)
}

/*
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
*/
