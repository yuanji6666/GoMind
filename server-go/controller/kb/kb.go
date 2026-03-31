package kb

import (
	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/common/code"
	"github.com/yuanji6666/gopherAI/controller"
	"github.com/yuanji6666/gopherAI/service/kb"
	"github.com/yuanji6666/gopherAI/schema"
)

type(
	GetKnowledgeBaseListRequest struct {
		Username string `json:"username" binding:"required"`
	}
	GetKnowledgeBaseListResponse struct {
		controller.Response
		KnowledgeBaseList []schema.KnowledgeBaseInfo `json:"knowledge_base_list"`
	}
	CreateKnowledgeBaseRequest struct {
		Username string `json:"username" binding:"required"`
		KBname string `json:"kb_name" binding:"required"`
	}
	CreateKnowledgeBaseResponse struct {
		controller.Response
		schema.KnowledgeBaseInfo `json:"knowledge_base_info,omitempty"`
	}
	
)



func GetKnowledgeBaseList(c *gin.Context) {
	res := new(GetKnowledgeBaseListResponse)

	username := c.GetString("username")
	if username == "" {
		c.JSON(200, res.CodeOf(code.CodeNotLogin))
		return
	}

	kbList, err := kb.GetKnowledgeBaseInfoList(username)
	if err != code.CodeSuccess{
		c.JSON(200, res.CodeOf(err))
		return
	}
	
	res.KnowledgeBaseList = kbList
	res.Success()
	c.JSON(200, res)
}

func CreateKnowledgeBase(c *gin.Context) {
	req := new(CreateKnowledgeBaseRequest)
	res := new(CreateKnowledgeBaseResponse)

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(200, res.CodeOf(code.CodeInvalidParams))
		return
	}

	kbInfo, err := kb.CreateKnowledgeBase(req.Username, req.KBname)
	if err != code.CodeSuccess {
		c.JSON(200, res.CodeOf(err))
		return
	}

	res.KnowledgeBaseInfo = kbInfo
	res.Success()
	c.JSON(200, res)

}