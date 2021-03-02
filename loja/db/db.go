package db

import (
	"database/sql"

	_ "github.com/lib/pq" // o _ informa ao compilador que essa biblioteca e usada em runtime
)

// Connect conecta ao banco de dados e retorna a conexao
func Connect() *sql.DB {

	connStr := "host=localhost dbname=jonnes_loja user=postgres password=surflife sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	}

	return db
}
