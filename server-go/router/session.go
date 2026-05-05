package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/domain/session"
)

func RegisterSessionRouter(r *gin.RouterGroup) {
	r.GET("/session/list", session.GetUserSessionsByUserName)
	r.POST("/session/create", session.CreateNewSessionAndSendMessageHandler)
	r.POST("/session/send", session.SendMessageHandler)
	r.POST("/session/history", session.GetHistoryBySessionIDWithIDHandler)
}
