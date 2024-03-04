package web

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"go-study/webook/internal/domain"
	"go-study/webook/internal/service"
	"net/http"
)

const (
	emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

// UserHandler 定义和用户有关的路由
type UserHandler struct {
	svc              *service.UserService
	emailRegexExp    *regexp.Regexp
	passwordRegexExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc:              svc,
		emailRegexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
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

	// Bind 方法会根据Content-Type指定的格式 解析数据到req里面
	// 解析错误 会直接返回400的错误
	if err := ctx.Bind(&req); err != nil {
		fmt.Println("错误")
		return
	}

	// 判断两次密码是否一致
	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "两次密码不一致")
		return
	}

	// 正则判断邮箱
	isEmail, err := u.emailRegexExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	if !isEmail {
		ctx.String(http.StatusOK, "邮箱不符合规范")
		return
	}

	// 正则判断密码
	isPassword, err := u.passwordRegexExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	if !isPassword {
		ctx.String(http.StatusOK, "密码必须包含数字、特殊字符，并且长度不能小于 8 位")
		return
	}

	err = u.svc.Signup(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	if err == service.ErrUserDuplicateEmail {
		ctx.String(http.StatusOK, "重复邮箱,请换一个邮箱")
		return
	}

	if err != nil {
		ctx.String(http.StatusOK, "服务器异常,注册失败")
		return
	}

	ctx.String(http.StatusOK, "注册成功")
}

// Login 用户登录
func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var loginReq LoginReq

	if err := ctx.Bind(&loginReq); err != nil {
		fmt.Println("错误")
		return
	}

}

// Edit 用户修改
func (u *UserHandler) Edit(ctx *gin.Context) {

}

// Profile 用户查看
func (u *UserHandler) Profile(ctx *gin.Context) {

}
