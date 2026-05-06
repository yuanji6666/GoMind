package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/domain/ai_session"
)

func RegisterSessionRouter(r *gin.RouterGroup) {
	r.GET("/ai_session/list", ai_session.GetUserSessionsByUserName)
	r.POST("/ai_session/create", ai_session.CreateNewSessionAndSendMessageHandler)
	r.POST("/ai_session/send", ai_session.SendMessageHandler)
	r.POST("/ai_session/history", ai_session.GetHistoryBySessionIDWithIDHandler)
}
