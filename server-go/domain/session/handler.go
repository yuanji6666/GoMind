package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/common/code"
	"github.com/yuanji6666/gopherAI/domain/message"
)

// Handler HTTP处理层

type (
	// GetUserSessionsByUserNameRequest 获取用户会话列表请求
	GetUserSessionsByUserNameRequest struct {
		UserName string `json:"username" binding:"required"`
	}
	// GetUserSessionsByUserNameResponse 获取用户会话列表响应
	GetUserSessionsByUserNameResponse struct {
		code.Response
		Sessions []SessionInfo `json:"sessions,omitempty"`
	}
	// CreateNewSessionAndSendMessageRequest 创建新会话并发送消息请求
	CreateNewSessionAndSendMessageRequest struct {
		UserName     string `json:"username" binding:"required"`
		UserQuestion string `json:"user_question" binding:"required"`
		UserKBID     string `json:"user_kb_id" binding:"required"`
	}
	// CreateNewSessionAndSendMessageResponse 创建新会话并发送消息响应
	CreateNewSessionAndSendMessageResponse struct {
		code.Response
		Answer    string `json:"answer"`
		SessionID string `json:"session_id"`
	}
	// SendMessageRequest 发送消息请求
	SendMessageRequest struct {
		SessionID    string `json:"session_id" binding:"required"`
		UserQuestion string `json:"user_question" binding:"required"`
	}
	// SendMessageResponse 发送消息响应
	SendMessageResponse struct {
		code.Response
		Answer string `json:"answer"`
	}
	// GetHistoryBySessionIDWithIDRequest 获取会话历史请求
	GetHistoryBySessionIDWithIDRequest struct {
		SessionID string `json:"session_id" binding:"required"`
		// last_id=0 表示从最早一条开始；required 会把数值 0 判为未填，故不能用 required
		LastID int64 `json:"last_id" binding:"gte=0"`
		Limit  int   `json:"limit" binding:"required,gte=1"`
	}
	// GetHistoryBySessionIDWithIDResponse 获取会话历史响应
	GetHistoryBySessionIDWithIDResponse struct {
		code.Response
		History []message.History `json:"history"`
	}
)

// GetUserSessionsByUserName 获取用户会话列表
func GetUserSessionsByUserName(c *gin.Context) {
	var res GetUserSessionsByUserNameResponse

	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeNotLogin))
		return
	}

	sessions, err := GetSessionByUsername(username)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	var sessionInfos []SessionInfo
	for _, session := range sessions {
		sessionInfos = append(sessionInfos, SessionInfo{
			Title:     session.Title,
			SessionID: session.ID,
			UserKBID:  session.UserKBID,
		})
	}
	res.Sessions = sessionInfos
	res.Success()
	c.JSON(http.StatusOK, res)
}

// CreateNewSessionAndSendMessage 创建新会话并发送消息
func CreateNewSessionAndSendMessageHandler(c *gin.Context) {
	var req CreateNewSessionAndSendMessageRequest
	var res CreateNewSessionAndSendMessageResponse
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, code.CodeInvalidParams)
		return
	}

	sessionID, answer, err := CreateNewSessionAndSendMessage(req.UserName, req.UserQuestion, req.UserKBID)
	if err != code.CodeSuccess {
		c.JSON(http.StatusOK, res.CodeOf(err))
		return
	}

	res.Answer = answer
	res.SessionID = sessionID
	res.Success()
	c.JSON(http.StatusOK, res)
}

// SendMessageHandler 发送消息
func SendMessageHandler(c *gin.Context) {
	var req SendMessageRequest
	var res SendMessageResponse
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, code.CodeInvalidParams)
		return
	}
	answer, err := SendMessage(req.UserQuestion, req.SessionID)
	if err != code.CodeSuccess {
		c.JSON(http.StatusOK, res.CodeOf(err))
		return
	}
	res.Answer = answer
	res.Success()
	c.JSON(http.StatusOK, res)
}

// GetHistoryBySessionIDWithIDHandler 获取会话历史
func GetHistoryBySessionIDWithIDHandler(c *gin.Context) {
	var req GetHistoryBySessionIDWithIDRequest
	var res GetHistoryBySessionIDWithIDResponse
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	history, err := message.GetHistoryBySessionIDWithID(req.SessionID, req.LastID, req.Limit)
	if err != code.CodeSuccess {
		c.JSON(http.StatusOK, res.CodeOf(err))
		return
	}

	res.History = history
	res.Success()
	c.JSON(http.StatusOK, res)
}
