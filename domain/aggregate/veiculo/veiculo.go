package veiculo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wellingtonpires/fase3-automoveis/domain/entity/veiculo"
	"github.com/wellingtonpires/fase3-automoveis/infrastructure/persistence"
)

func Cadastro(c *gin.Context) {

	token := c.GetHeader("Authorization")
	if token == "" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"erro": "Token não informado"})

	} else if token == "aloalorapaziada" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"erro": "Token não informado"})
	} else {
		var ve veiculo.Veiculo
		err := c.ShouldBindJSON(&ve)
		if err != nil {
			fmt.Println(err.Error())
		}
		persistence.CadastraVeiculo(ve)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"resultado": "Erro: " + err.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, gin.H{"resultado": "Veiculo cadastrado na base"})
		}
	}

}

func Atualizacao(c *gin.Context) {
	var ve veiculo.Veiculo
	err := c.ShouldBindJSON(&ve)
	if err != nil {
		fmt.Println(err.Error())
	}
	persistence.AtualizaVeiculo(ve)
	c.IndentedJSON(http.StatusOK, gin.H{"resultado": "Veículo atualizado na base"})
}

func Exclusao(c *gin.Context) {
	var ve veiculo.Veiculo
	err := c.ShouldBindJSON(&ve)
	if err != nil {
		fmt.Println(err.Error())
	}
	persistence.ExcluiVeiculo(ve)
	c.IndentedJSON(http.StatusOK, gin.H{"resultado": "Veículo deletado"})
}

func ConsultaPorPreco(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persistence.ConsultaPorPreco())
}

func ConsultaVendidos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persistence.ConsultaVendidos())
}
