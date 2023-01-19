package controller

import "github.com/gin-gonic/gin"

func HelloWord (ctx *gin.Context) {
	ctx.JSON(200, "Hello word")
}