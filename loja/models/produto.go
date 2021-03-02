package models

import "github.com/jonnes-nascimento/db"

// Produto e a estrutura que contem as caracteristicas do produto da loja
type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// ProdutoFindAll retorna todos os produtos cadastrados no BD
func ProdutoFindAll() []Produto {

	db := db.Connect()
	defer db.Close()

	rsProdutos, err := db.Query("select * from produto order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for rsProdutos.Next() {

		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := rsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	return produtos
}

// ProdutoFindByID retorna o produto com o id passado no parametro do BD
func ProdutoFindByID(id string) Produto {

	db := db.Connect()
	defer db.Close()

	produtoRs, err := db.Query("select id, nome, descricao, preco, quantidade from produto where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for produtoRs.Next() {

		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := produtoRs.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.ID = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}

	return produto
}

// ProdutoInsert cadastra um novo produto no BD
func ProdutoInsert(nome, descricao string, preco float64, quantidade int) {

	db := db.Connect()
	defer db.Close()

	insertProduto, err := db.Prepare("insert into produto (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertProduto.Exec(nome, descricao, preco, quantidade)
}

// ProdutoDelete apaga um produto cadastrado no BD
func ProdutoDelete(id string) {

	db := db.Connect()
	defer db.Close()

	deleteProduto, err := db.Prepare("delete from produto where id = $1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduto.Exec(id)
}

// ProdutoUpdate altera as informacos do produto no BD
func ProdutoUpdate(id int, nome, descricao string, preco float64, quantidade int) {

	db := db.Connect()
	defer db.Close()

	updateProduto, err := db.Prepare("update produto set nome = $1, descricao = $2, preco = $3, quantidade = $4 where id = $5")
	if err != nil {
		panic(err.Error())
	}

	updateProduto.Exec(nome, descricao, preco, quantidade, id)
}
