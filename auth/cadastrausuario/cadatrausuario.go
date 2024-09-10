package cadastrausuario

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
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

type Usuario struct {
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
	Email   string `json:"email"`
	Cpf     string `json:"cpf"`
}

func CadastraUsuario(c *gin.Context) {
	var u Usuario
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Conectado DB")
	}
	defer con.Close()

	rows, err := con.Query(`INSERT INTO user VALUES ($1, $2, $3, $4)`, u.Usuario, u.Senha, u.Email, u.Cpf)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()
}
