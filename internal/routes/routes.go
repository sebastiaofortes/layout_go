package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastiaofortes/layout_go/internal/controller"
	"github.com/sebastiaofortes/layout_go/internal/service"
)

func ConfigureRoutes(r *gin.Engine) *gin.Engine {
	s := service.NewDefaultService()
	e := controller.NewExpectativaController(s)

	r.GET("/expectativa-pais", e.ExpectativaDeVidaPorPais())
	r.GET("/expectativa-idade", e.ExpectativaDeVidaPorIdade())
	r.GET("/hello", controller.HelloWord)
	return r
}
