package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	Routes(router)
	router.Run(":8081")
}

func Routes(route *gin.Engine) {
	v := route.Group("/external")
	v.GET("/consulta-por-preco", func(c *gin.Context) {
		// Defina a URL que você deseja consultar
		url := "http://localhost:8080/veiculo/consulta-por-preco"

		// Crie a requisição GET
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Erro ao fazer a requisição: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao obter dados externos"})
			return
		}
		defer resp.Body.Close()

		// Leia o corpo da resposta
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Erro ao ler o corpo da resposta: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao ler dados externos"})
			return
		}

		// Adiciona o conteúdo da resposta no contexto do Gin
		c.JSON(http.StatusOK, gin.H{
			"data": string(body),
		})
	})
}
