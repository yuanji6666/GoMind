package kb

import (
	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/common/code"
)

// Handler HTTP处理层

type (
	GetKnowledgeBaseListRequest struct {
		Username string `json:"username" binding:"required"`
	}
	GetKnowledgeBaseListResponse struct {
		code.Response
		KnowledgeBaseList []KnowledgeBaseInfo `json:"knowledge_base_list"`
	}
	CreateKnowledgeBaseRequest struct {
		Username string `json:"username" binding:"required"`
		KBname   string `json:"kb_name" binding:"required"`
	}
	CreateKnowledgeBaseResponse struct {
		code.Response
		KnowledgeBaseInfo `json:"knowledge_base_info,omitempty"`
	}
)

// GetKnowledgeBaseList 获取用户知识库列表
func GetKnowledgeBaseList(c *gin.Context) {
	res := new(GetKnowledgeBaseListResponse)

	username := c.GetString("username")
	if username == "" {
		c.JSON(200, res.CodeOf(code.CodeNotLogin))
		return
	}

	kbList, err := GetKnowledgeBaseInfoList(username)
	if err != code.CodeSuccess {
		c.JSON(200, res.CodeOf(err))
		return
	}

	res.KnowledgeBaseList = kbList
	res.Success()
	c.JSON(200, res)
}

// CreateKnowledgeBase 创建新的知识库
func CreateNewKnowledgeBase(c *gin.Context) {
	req := new(CreateKnowledgeBaseRequest)
	res := new(CreateKnowledgeBaseResponse)

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(200, res.CodeOf(code.CodeInvalidParams))
		return
	}

	kbInfo, err := CreateKB(req.Username, req.KBname)
	if err != code.CodeSuccess {
		c.JSON(200, res.CodeOf(err))
		return
	}

	res.KnowledgeBaseInfo = kbInfo
	res.Success()
	c.JSON(200, res)
}
