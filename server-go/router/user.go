package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/domain/user"
)

func RegisterUserRouter(r *gin.RouterGroup) {
	{
		r.POST("/register", user.HandleRegister)
		r.POST("/captcha", user.HandleCaptcha)
		r.POST("/login", user.HandleLogin)
	}
}
