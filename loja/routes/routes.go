package routes

import (
	"net/http"

	"github.com/jonnes-nascimento/controllers"
)

// LoadRoutes carrega todos os mapeamentos de rota
func LoadRoutes() {

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	http.HandleFunc("/", controllers.ProdutoListar)
	http.HandleFunc("/incluirProduto", controllers.ProdutoIncluir)
	http.HandleFunc("/alterarProduto", controllers.ProdutoAlterar)
	http.HandleFunc("/produtoInsert", controllers.ProdutoInsert)
	http.HandleFunc("/produtoDelete", controllers.ProdutoDelete)
	http.HandleFunc("/produtoUpdate", controllers.ProdutoUpdate)
}
