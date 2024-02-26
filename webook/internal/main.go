package main

import (
	"github.com/gin-gonic/gin"
	"go-study/webook/internal/web"
)

func main() {

	server := gin.Default()

	u := &web.UserHandler{}
	u.RegisterUserRoutes(server)

	server.Run(":8080")
}
