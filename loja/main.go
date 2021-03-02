package main

import (
	"fmt"
	"net/http"

	"github.com/jonnes-nascimento/go/loja/routes"
)

func main() {

	println("Hello!")
	fmt.Println("'Loja do Jonnes' está executando e está disponível em localhost:8000")

	routes.LoadRoutes()

	http.ListenAndServe(":8000", nil)
}
