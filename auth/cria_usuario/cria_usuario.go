package cria_usuario

import (
	"database/sql"
	"fmt"
)

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://postgres:123@postgres:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Conectado ao banco!")
	}

	err = db.Ping()

	return db, err
}

func cadastraCliente(usuario string, senha string, email string, cpf string) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Conectado DB")
	}
	defer con.Close()

	rows, err := con.Query(`INSERT INTO user VALUES ($1, $2, $3, $4)`, usuario, senha, email, cpf)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()
}
