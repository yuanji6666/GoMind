package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/domain/kb"
)

func RegisterKBRouter(r *gin.RouterGroup) {
	r.GET("/kb/list", kb.GetKnowledgeBaseList)
	r.POST("/kb/create", kb.CreateNewKnowledgeBase)
}
