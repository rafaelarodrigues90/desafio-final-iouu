package main

import (
	"desafio-final-iouu/routes"
)

func main() {

	// iniciando echo
	e := routes.Index()

	// iniciando servidor na porta 8000
	e.Logger.Fatal(e.Start(":8000"))
}
