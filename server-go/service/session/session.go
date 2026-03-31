package session

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/yuanji6666/gopherAI/common/code"
	"github.com/yuanji6666/gopherAI/config"
	"github.com/yuanji6666/gopherAI/dao/session"
	"github.com/yuanji6666/gopherAI/schema"
	"resty.dev/v3"
)

type ChatResponse struct {
	Answer string `json:"answer"`
}

// CreateSessionAndSendMessage 创建新会话并发送消息，返回会话ID，第一条回复，状态码
func CreateNewSessionAndSendMessage(userName, userQuestion, userKBID string) (string, string, code.Code) {
	// 新建会话存储在数据库
	newSession := &schema.Session{
		ID:       uuid.New().String(),
		UserName: userName,
		UserKBID: userKBID,
		Title:    userQuestion,
	}

	_, err := session.CreateSession(newSession)

	if err != nil {
		log.Println("CreateSessionAndSendMessage CreateSession error:", err)
		return "", "", code.CodeServerBusy
	}

	// 调用远程服务获取回复
	client := resty.New()
	defer client.Close()

	answer := &ChatResponse{}
	url := fmt.Sprintf("http://%s:%s", config.GetConfig().RAGConfig.Host, config.GetConfig().RAGConfig.Port)
	resp, err := client.R().
		SetBody(map[string]interface{}{
			"message": userQuestion,
			"history": []map[string]string{},
			"top_k": 4,
		}).
		SetResult(answer).
		Post(url+"/knowledge-bases/"+userKBID+"/chat")
	if err != nil {
		log.Println("RAG request error:", err)
		return newSession.ID, "", code.CodeServerBusy
	}
	if !resp.IsSuccess() {
		log.Printf("RAG request failed: status=%d body=%s", resp.StatusCode(), resp.String())
		return newSession.ID, "", code.CodeServerBusy
	}
	return newSession.ID, answer.Answer, code.CodeSuccess
}

func SendMessage(userQuestion, SessionID string) (string, code.Code) {
	session, err:= session.GetSessionByID(SessionID)
	if err != nil {
		log.Println("GetSessionByID error:", err)
		return "", code.CodeServerBusy
	}
	userKBID := session.UserKBID
	url := fmt.Sprintf("http://%s:%s", config.GetConfig().RAGConfig.Host, config.GetConfig().RAGConfig.Port)
	client := resty.New()
	defer client.Close()
	answer := &ChatResponse{}
	resp, err := client.R().
		SetBody(map[string]interface{}{
			"message": userQuestion,
			"history": []map[string]string{},
			"top_k": 4,
		}).
		SetResult(answer).
		Post(url+"/knowledge-bases/"+userKBID+"/chat")
	
	if err != nil {
		log.Println("RAG request error:", err)
		return "", code.CodeServerBusy
	}
	if !resp.IsSuccess() {
		log.Printf("RAG request failed: status=%d body=%s", resp.StatusCode(), resp.String())
		return "", code.CodeServerBusy
	}
	return answer.Answer, code.CodeSuccess
}
