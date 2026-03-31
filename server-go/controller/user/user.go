package user

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/common/code"
	"github.com/yuanji6666/gopherAI/controller"
	"github.com/yuanji6666/gopherAI/service/user"
)

type(
	//前后端接口格式
	RegisterRequest struct{
		Email		string`json:"email" binding:"required"`
		Captcha		string`json:"captcha"`
		Password	string`json:"password"`
	}
	RegisterResponse struct{
		controller.Response
		Token		string`json:"token,omitempty"`
	}
	LoginRequest	struct{
		Username	string`json:"username"`
		Password	string`json:"password"`
	}
	LoginResponse	struct{
		controller.Response
		Token		string`json:"token,omitempty"`
	}
	CaptchaRequest	struct{
		Email		string`json:"email" binding:"required"`
	}
	CaptchaResponse	struct{
		controller.Response
	}
)

func Register(ctx *gin.Context){
	req := new(RegisterRequest)
	res := new(RegisterResponse)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	//交给service层处理
	token, code_ := user.Register(req.Email, req.Password, req.Captcha)

	if code_ != code.CodeSuccess {
		ctx.JSON(http.StatusOK, res.CodeOf(code_))
		return
	}

	res.Success()
	res.Token = token
	ctx.JSON(http.StatusOK, res)

}

func Login(ctx *gin.Context){
	req := new(LoginRequest)
	res := new(LoginResponse)

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	token, code_ := user.Login(req.Username, req.Password)

	if code_ != code.CodeSuccess {
		ctx.JSON(http.StatusOK, res.CodeOf(code_))
		return
	}

	res.Token = token
	res.Success()

	ctx.JSON(http.StatusOK, res)

}



func HandleCaptcha(ctx *gin.Context){
	req := new(CaptchaRequest)
	res := new(CaptchaResponse)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	//交给service层处理
	code_ := user.SendCaptcha(req.Email)

	if code_ != code.CodeSuccess {
		ctx.JSON(http.StatusOK, res.CodeOf(code_))
		return
	}

	res.Success()
	ctx.JSON(http.StatusOK, res)

}