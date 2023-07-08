package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastiaofortes/layout_go/internal/controller"
	"github.com/sebastiaofortes/layout_go/internal/platform/db"
	"github.com/sebastiaofortes/layout_go/internal/platform/repository"
	"github.com/sebastiaofortes/layout_go/internal/service"
)

func ConfigureRoutes(r *gin.Engine) *gin.Engine {
	paisR := repository.NewImplementsPaisRepository()
	pessoaR := repository.NewImplementsPessoaRepository()

	s := service.NewDefaultService(pessoaR, paisR)
	e := controller.NewExpectativaController(s)

	db.Get()

	r.GET("/expectativa-pais", e.ExpectativaDeVidaPorPais())
	r.GET("/expectativa-idade", e.ExpectativaDeVidaPorIdade())
	r.GET("/hello", controller.HelloWord)
	return r
}
