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

func (c *ExpectativaController) ExpectativaDeVidaPorPais(ctx *gin.Context) {
	ctx.JSON(200, "expectativa por país")
}

func (c *ExpectativaController) ExpectativaDeVidaPorIdade(ctx *gin.Context) {
	ctx.JSON(200, "expectativa por país")
}