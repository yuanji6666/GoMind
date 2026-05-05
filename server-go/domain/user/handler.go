package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/common/code"
)

// Handler HTTP处理层

type (
	// RegisterRequest 注册请求
	RegisterRequest struct {
		Email    string `json:"email" binding:"required"`
		Captcha  string `json:"captcha"`
		Password string `json:"password"`
	}
	// RegisterResponse 注册响应
	RegisterResponse struct {
		code.Response
		Token string `json:"token,omitempty"`
	}
	// LoginRequest 登录请求
	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	// LoginResponse 登录响应
	LoginResponse struct {
		code.Response
		Token string `json:"token,omitempty"`
	}
	// CaptchaRequest 验证码请求
	CaptchaRequest struct {
		Email string `json:"email" binding:"required"`
	}
	// CaptchaResponse 验证码响应
	CaptchaResponse struct {
		code.Response
	}
)

// HandleRegister 处理用户注册
func HandleRegister(ctx *gin.Context) {
	req := new(RegisterRequest)
	res := new(RegisterResponse)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 交给 service 层处理
	token, code_ := Register(req.Email, req.Password, req.Captcha)

	if code_ != code.CodeSuccess {
		ctx.JSON(http.StatusOK, res.CodeOf(code_))
		return
	}

	res.Success()
	res.Token = token
	ctx.JSON(http.StatusOK, res)
}

// HandleLogin 处理用户登录
func HandleLogin(ctx *gin.Context) {
	req := new(LoginRequest)
	res := new(LoginResponse)

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	token, code_ := Login(req.Username, req.Password)

	if code_ != code.CodeSuccess {
		ctx.JSON(http.StatusOK, res.CodeOf(code_))
		return
	}

	res.Token = token
	res.Success()

	ctx.JSON(http.StatusOK, res)
}

// HandleCaptcha 处理验证码请求
func HandleCaptcha(ctx *gin.Context) {
	req := new(CaptchaRequest)
	res := new(CaptchaResponse)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 交给 service 层处理
	code_ := SendCaptcha(req.Email)

	if code_ != code.CodeSuccess {
		ctx.JSON(http.StatusOK, res.CodeOf(code_))
		return
	}

	res.Success()
	ctx.JSON(http.StatusOK, res)
}
