package main

import (
	"fmt"
	"net/http"

	"github.com/jonnes-nascimento/loja/routes"
)

func main() {

	fmt.Println("'Loja do Jonnes' executando na porta 8000")

	routes.LoadRoutes()

	http.ListenAndServe(":8000", nil)
}
