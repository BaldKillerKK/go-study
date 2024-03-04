package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-study/webook/internal/repository"
	"go-study/webook/internal/repository/dao"
	"go-study/webook/internal/service"
	"go-study/webook/internal/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {

	server := gin.Default()
	db, err := gorm.Open(mysql.Open("root:2wsx#EDC@tcp(192.168.132.20:3306)/webbook"))
	if err != nil {
		panic(err.Error())
	}
	userDao := dao.NewUserDao(db)
	userRepository := repository.NewUserRepository(userDao)
	userService := service.NewUserService(userRepository)
	userHandler := web.NewUserHandler(userService)
	userHandler.RegisterUserRoutes(server)

	// 使用middleware 解决跨域问题
	server.Use(func(context *gin.Context) {
		fmt.Println("第一个middlerware")
	})

	server.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return false
		},
		MaxAge: 12 * time.Hour,
	}))
	server.Run(":8080")
}
