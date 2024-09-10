package valida_token

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func verificaCadastro(usuario string, senha string) (status string) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Conectado DB")
	}
	defer con.Close()

	rows, err := con.Query(`SELECT * FROM user WHERE user = $1, password = $2`, usuario, senha)
	if err != nil {
		fmt.Println(err.Error())
		return "FALHA"
	}
	defer rows.Close()
	return "SUCESSO"
}

var secretKey = []byte("fase3sub")

func createToken(username string, password string) (string, error) {
	if verificaCadastro(username, password) == "teste" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{
				"username": username,
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
			})

		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			return "", err
		}

		return tokenString, nil
	} else {
		return "Cliente não existe na base.", nil
	}

}
