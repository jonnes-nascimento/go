package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/jonnes-nascimento/go/loja/models"
)

//var t = template.Must(template.ParseGlob("templates/*.html"))
var t = template.Must(template.ParseFiles(
	"templates/_head.html",
	"templates/_navbar.html",
	"templates/listarProdutos.html",
	"templates/incluirProduto.html",
	"templates/alterarProduto.html",
))

// ProdutoListar busca todos os produtos cadastrados no BD e exibe no template listarProdutos
func ProdutoListar(w http.ResponseWriter, r *http.Request) {

	produtos := models.ProdutoFindAll()

	t.ExecuteTemplate(w, "ListarProdutos", produtos)
}

// ProdutoIncluir carrega a tela para cadastramento de um novo produto
func ProdutoIncluir(w http.ResponseWriter, r *http.Request) {

	t.ExecuteTemplate(w, "IncluirProduto", nil)
}

// ProdutoAlterar carrega a tela para alteracao de um produto
func ProdutoAlterar(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id")

	produto := models.ProdutoFindByID(idProduto)

	t.ExecuteTemplate(w, "AlterarProduto", produto)
}

// ProdutoInsert captura os dados do produto no formulario e repassa ao modelo para inclusao no BD
func ProdutoInsert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Ocorreu um erro na conversão do preço")
		}

		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Ocorreu um erro na conversão da quantidade")
		}

		models.ProdutoInsert(nome, descricao, precoConv, quantidadeConv)
	}

	http.Redirect(w, r, "/", 301)
}

// ProdutoDelete captura o id do produto no formulario e repassa ao modelo para exclusao no BD
func ProdutoDelete(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id") // nao precisa converter por o id vai ser usado para montar a query (string)

	models.ProdutoDelete(idProduto)

	http.Redirect(w, r, "/", 301)
}

// ProdutoUpdate captura as informacoes do produto do formulario e repassa ao modelo para alterar no BD
func ProdutoUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Ocorreu um erro na conversão do id")
		}

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Ocorreu um erro na conversão do preço")
		}

		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Ocorreu um erro na conversão da quantidade")
		}

		models.ProdutoUpdate(idConv, nome, descricao, precoConv, quantidadeConv)

		http.Redirect(w, r, "/", 301)
	}
}
