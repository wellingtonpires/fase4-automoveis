package criatoken

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func verificaCadastro(login string, senha string) (status bool) {
	fmt.Print(login + senha)
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Conectado DB")
	}
	defer con.Close()
	rows, err := con.Query(`SELECT * FROM usuario WHERE login = $1 AND password = $2`, login, senha)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer rows.Close()
	return true
}

type Usuario struct {
	Login string `json:"login"`
	Senha string `json:"senha"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
	Admin bool   `json:"admin"`
}

func consultaRole(login string, senha string) (role string) {
	var admin bool
	fmt.Print(login + senha)
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Conectado DB")
	}
	defer con.Close()
	rows, err := con.Query(`SELECT * FROM usuario WHERE login = $1 AND password = $2`, login, senha)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var u Usuario
		rows.Scan(u.Login, u.Senha, u.Email, u.Cpf, u.Admin)
		admin = u.Admin
	}
	defer rows.Close()
	if admin {
		return "admin"
	} else {
		return "user"
	}
}

var secretKey = []byte("fase3sub")

func CriaToken(c *gin.Context) {
	if verificaCadastro(c.Query("login"), c.Query("senha")) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{
				"username": c.Query("login"),
				"role":     consultaRole(c.Query("login"), c.Query("senha")),
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
			})
		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao assinar token"})
		} else {
			c.IndentedJSON(http.StatusCreated, gin.H{"token": tokenString})
		}
	} else {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"erro": "Usuário não existe na base"})
	}
}
