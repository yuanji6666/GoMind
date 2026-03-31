package router

import(
	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/controller/kb"
)

func RegisterKBRouter(r *gin.RouterGroup) {
	r.GET("/kb/list", kb.GetKnowledgeBaseList)
	r.POST("/kb/create", kb.CreateKnowledgeBase)
}