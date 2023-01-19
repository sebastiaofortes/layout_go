package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastiaofortes/layout_go/internal/routes"
)

func main() {
	r := gin.Default()
	r = routes.ConfigureRoutes(r)
	r.Run()
}
