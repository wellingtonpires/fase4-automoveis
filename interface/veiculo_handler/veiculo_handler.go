package veiculo_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wellingtonpires/fase3-automoveis/domain/aggregate/veiculo"
)

func Routes(route *gin.Engine) {
	v := route.Group("/veiculo")
	v.GET("/vendidos", veiculo.ConsultaVendidos)
	v.GET("/consulta-por-preco", veiculo.ConsultaPorPreco)
	v.POST("/cadastra-veiculo", veiculo.Cadastro)
	v.PATCH("/atualiza-veiculo", veiculo.Atualizacao)
	v.DELETE("/exclui-veiculo", veiculo.Exclusao)

	u := route.Group("/user")
	u.POST("/cadastra-cliente", cria_token.criaToken)
	u.POST("/cria-token", veiculo.Cadastro)
	u.GET("/valida-token", veiculo.Cadastro)
}
