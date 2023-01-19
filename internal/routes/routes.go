package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastiaofortes/layout_go/internal/controller"
)

func ConfigureRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/hello", controller.HelloWord)
	return r
}