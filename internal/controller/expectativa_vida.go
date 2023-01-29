package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastiaofortes/layout_go/internal/service"
)

type ExpectativaController struct{
	s service.EsperançaDeVida
}

func NewExpectativaController(s service.EsperançaDeVida) ExpectativaController {
	c := ExpectativaController{s: s}
	return c
}

//quando estamos usando o framework gin
//recomenda-se que as funções de handler sejam geradas por outra função (gin.HandlerFunc)
//pois assim podemos passar parametros adicionais e injetar dependencias nas funções de handler
func (c *ExpectativaController) ExpectativaDeVidaPorPais() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, "expectativa por país")
	}
}

func (c *ExpectativaController) ExpectativaDeVidaPorIdade() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, "expectativa por idade")
	}
}
