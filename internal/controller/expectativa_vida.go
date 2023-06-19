package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExpectativaController struct {
	s EsperançaDeVida
}

func NewExpectativaController(s EsperançaDeVida) ExpectativaController {
	c := ExpectativaController{s: s}
	return c
}

// quando estamos usando o framework gin
// recomenda-se que as funções de handler sejam geradas por outra função (gin.HandlerFunc)
// pois assim podemos passar parametros adicionais e injetar dependencias nas funções de handler
func (c *ExpectativaController) ExpectativaDeVidaPorPais() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Query("pais"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "ivalid pais id")
			return
		}
		r, err := c.s.CalcularExpectativaDeVidaPorPais(int32(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, "not results found")
			return
		}
		str := fmt.Sprintf("%.3f", r)
		ctx.JSON(200, str)
	}
}

func (c *ExpectativaController) ExpectativaDeVidaPorIdade() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Query("idade"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "ivalid pais id")
			return
		}
		r, err := c.s.CalcularExpectativaDeVidaPorIdade(int32(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, "not results found")
			return
		}
		ctx.JSON(200, r)
	}
}
