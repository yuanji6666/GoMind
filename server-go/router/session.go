package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/controller/session"
)

func RegisterSessionRouter(r * gin.RouterGroup) {
	r.GET("/session/list", session.GetUserSessionsByUserName)
	r.POST("/session/create", session.CreateNewSessionAndSendMessage)
	r.POST("/session/send", session.SendMessage)
	r.POST("/session/history", session.GetHistoryBySessionIDWithID)
}