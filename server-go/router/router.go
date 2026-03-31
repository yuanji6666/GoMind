package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/middleware/auth"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	

	enterRouter := r.Group("/api/v1")
	{
		//注册用户注册，登陆，验证码路由
		RegisterUserRouter(enterRouter.Group("/user"))
	}
	{
		AIGroup := enterRouter.Group("/AI")
		AIGroup.Use(auth.Auth())
		RegisterKBRouter(AIGroup)
		RegisterSessionRouter(AIGroup)
	}

	
	return r
}