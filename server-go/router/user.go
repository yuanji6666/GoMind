package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/controller/user"
)

func RegisterUserRouter(r *gin.RouterGroup){
	{
		r.POST("/register", user.Register)
		r.POST("/captcha", user.HandleCaptcha)
		r.POST("/login",user.Login )
	}
}