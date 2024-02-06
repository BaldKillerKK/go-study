package web

import "github.com/gin-gonic/gin"

// UserHandler 定义和用户有关的路由
type UserHandler struct {
}

// RegisterUserRoutes 用户路由注册
func (u *UserHandler) RegisterUserRoutes(server *gin.Engine) {
	// 路由分组
	ug := server.Group("/users")
	ug.GET("/profile", u.Profile)
	ug.POST("/login", u.Login)
	ug.POST("/signup", u.Signup)
	ug.POST("/edit", u.Edit)
}

// Signup 用户注册
func (u *UserHandler) Signup(ctx *gin.Context) {
	type signupReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req signupReq

	if err := ctx.Bind(&req); err != nil {
		return
	}

	ctx.String(200, "signup %v", req)
}

// Login 用户登录
func (u *UserHandler) Login(ctx *gin.Context) {

}

// Edit 用户修改
func (u *UserHandler) Edit(ctx *gin.Context) {

}

// Profile 用户查看
func (u *UserHandler) Profile(ctx *gin.Context) {

}
