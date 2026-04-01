package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/common/code"
	"github.com/yuanji6666/gopherAI/controller"
	"github.com/yuanji6666/gopherAI/schema"
	daoSession "github.com/yuanji6666/gopherAI/dao/session"
	serviceSession "github.com/yuanji6666/gopherAI/service/session"
	serviceMessage "github.com/yuanji6666/gopherAI/service/message"


)
type(
	GetUserSessionsByUserNameRequest struct{
		UserName string `json:"username" binding:"required"`
	}
	GetUserSessionsByUserNameResponse struct{
		controller.Response
		Sessions []schema.SessionInfo `json:"sessions,omitempty"`
	}
	CreateNewSessionAndSendMessageRequest struct{
		UserName string `json:"username" binding:"required"`
		UserQuestion string `json:"user_question" binding:"required"`
		UserKBID string `json:"user_kb_id" binding:"required"`
	}
	CreateNewSessionAndSendMessageResponse struct{
		controller.Response
		Answer string `json:"answer"`
		SessionID string `json:"session_id"`
	}
	SendMessageRequest struct{
		SessionID string `json:"session_id" binding:"required"`
		UserQuestion string `json:"user_question" binding:"required"`
	}
	SendMessageResponse struct{
		controller.Response
		Answer string `json:"answer"`
	}
	GetHistoryBySessionIDWithIDRequest struct{
		SessionID string `json:"session_id" binding:"required"`
		// last_id=0 表示从最早一条开始；required 会把数值 0 判为未填，故不能用 required
		LastID int64 `json:"last_id" binding:"gte=0"`
		Limit  int   `json:"limit" binding:"required,gte=1"`
	}
	GetHistoryBySessionIDWithIDResponse struct{
		controller.Response
		History []schema.History `json:"history"`
	}
)

func GetUserSessionsByUserName(c *gin.Context) {
	var res GetUserSessionsByUserNameResponse

	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeNotLogin))
		return
	}

	sessions, err := daoSession.GetSessionByUsername(username)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	var sessionInfos []schema.SessionInfo
	for _, session := range sessions {
		sessionInfos = append(sessionInfos, schema.SessionInfo{
			Title: session.Title,
			SessionID: session.ID,
			UserKBID: session.UserKBID,
		})
	}
	res.Sessions = sessionInfos
	res.Success()
	c.JSON(http.StatusOK, res)
}

func CreateNewSessionAndSendMessage(c *gin.Context) {
	var req CreateNewSessionAndSendMessageRequest
	var res CreateNewSessionAndSendMessageResponse
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, code.CodeInvalidParams)
		return
	}

	sessionID, answer, err := serviceSession.CreateNewSessionAndSendMessage(req.UserName, req.UserQuestion, req.UserKBID)
	if err != code.CodeSuccess {
		c.JSON(http.StatusOK, res.CodeOf(err))
		return
	}

	res.Answer = answer
	res.SessionID = sessionID
	res.Success()
	c.JSON(http.StatusOK, res)
}

func SendMessage(c *gin.Context) {
	var req SendMessageRequest
	var res SendMessageResponse
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, code.CodeInvalidParams)
		return
	}
	answer, err := serviceSession.SendMessage(req.UserQuestion, req.SessionID)
	if err != code.CodeSuccess {
		c.JSON(http.StatusOK, res.CodeOf(err))
		return
	}
	res.Answer = answer
	res.Success()
	c.JSON(http.StatusOK, res)
}

func GetHistoryBySessionIDWithID(c *gin.Context) {
	var req GetHistoryBySessionIDWithIDRequest
	var res GetHistoryBySessionIDWithIDResponse
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	history, err := serviceMessage.GetHistoryBySessionIDWithID(req.SessionID, req.LastID, req.Limit)
	if err != code.CodeSuccess {
		c.JSON(http.StatusOK, res.CodeOf(err))
		return
	}	

	res.History = history
	res.Success()
	c.JSON(http.StatusOK, res)
}