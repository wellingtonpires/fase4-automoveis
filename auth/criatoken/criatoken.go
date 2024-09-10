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

func verificaCadastro(usuario string, senha string) (status bool) {
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
		return false
	}
	defer rows.Close()
	return true
}

var secretKey = []byte("fase3sub")

func CriaToken(c *gin.Context) {
	if verificaCadastro(c.Query("usuario"), c.Query("senha")) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{
				"username": c.Query("usuario"),
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
