package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-study/webook/internal/web"
)

func main() {

	server := gin.Default()

	u := web.NewUserHandler()
	u.RegisterUserRoutes(server)

	// 使用middleware 解决跨域问题
	server.Use(func(context *gin.Context) {
		fmt.Println("第一个middlerware")
	})

	server.Run(":8080")
}
