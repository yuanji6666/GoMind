package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/common/code"
	"github.com/yuanji6666/gopherAI/controller"
	"github.com/yuanji6666/gopherAI/utils/myjwt"
)

// Auth 提供登陆校验中间件
// 解析token并把username加入context
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context){
		res := new(controller.Response)
		token := ctx.GetHeader("Authorization")
		if token != "" && strings.HasPrefix(token, "Bearer ") {
			token,_ = strings.CutPrefix(token, "Bearer ")
		} else {
			token = ctx.Query("token")
		}

		if token == "" {
			ctx.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidToken))
			ctx.Abort()
			return 
		}

		log.Println(token)

		username, ok := myjwt.ParseToken(token)

		if !ok {
			ctx.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidToken))
			ctx.Abort()
			return 
		}

		ctx.Set("username", username)
		ctx.Next()
	}
}