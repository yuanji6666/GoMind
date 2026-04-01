package session

import (
	"context"
	"fmt"
	"log"
	"time"

	"encoding/json"

	"github.com/google/uuid"
	"github.com/yuanji6666/gopherAI/common/code"
	"github.com/yuanji6666/gopherAI/common/redis"
	"github.com/yuanji6666/gopherAI/config"
	"github.com/yuanji6666/gopherAI/dao/message"
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
	
	// 消息持久化到redis
	key := "session:"+newSession.ID
	ctx := context.Background()
	
	userMsg, err := json.Marshal(schema.History{
		Role: "user",
		Content: userQuestion,
	})
	
	assistantMsg, err := json.Marshal(schema.History{
		Role: "assistant",
		Content: answer.Answer,
	})


	// 消息持久化到数据库
	go func() {
		pipe := redis.Rdb.Pipeline()
		pipe.RPush(ctx, key, userMsg, assistantMsg)
		pipe.LTrim(ctx, key, -10, -1)
		pipe.Expire(ctx, key, 20*time.Minute)
		pipe.Exec(ctx)
	}()

	go message.CreateMessage(&schema.Message{
		SessionID: newSession.ID,
		UserName: userName,
		Content: userQuestion,
		IsUser: true,
	})
	
	go message.CreateMessage(&schema.Message{
		SessionID: newSession.ID,
		UserName: userName,
		Content: answer.Answer,
		IsUser: false,
	})
	
	
	return newSession.ID, answer.Answer, code.CodeSuccess
}

// SendMessage 发送消息的业务逻辑
func SendMessage(userQuestion, SessionID string) (string, code.Code) {
	session, err:= session.GetSessionByID(SessionID)
	if err != nil {
		log.Println("GetSessionByID error:", err)
		return "", code.CodeServerBusy
	}
	
	// 获取历史消息
	historyMsgs, err := redis.Rdb.LRange(context.Background(), "session:"+SessionID, 0, -1).Result()
	if err != nil {
		log.Println("LRange error:", err)
		return "", code.CodeServerBusy
	}
	
	history := make([]map[string]string, len(historyMsgs))
	for i, msg := range historyMsgs {
		json.Unmarshal([]byte(msg), &history[i])
	}
	
	fmt.Println(history)

	
	url := fmt.Sprintf("http://%s:%s", config.GetConfig().RAGConfig.Host, config.GetConfig().RAGConfig.Port)
	client := resty.New()
	defer client.Close()
	answer := &ChatResponse{}
	resp, err := client.R().
		SetBody(map[string]interface{}{
			"message": userQuestion,
			"history": history,
			"top_k": 4,
		}).
		SetResult(answer).
		Post(url+"/knowledge-bases/"+session.UserKBID+"/chat")
	
	if err != nil {
		log.Println("RAG request error:", err)
		return "", code.CodeServerBusy
	}
	if !resp.IsSuccess() {
		log.Printf("RAG request failed: status=%d body=%s", resp.StatusCode(), resp.String())
		return "", code.CodeServerBusy
	}
	
	// 消息持久化到redis
	key := "session:"+SessionID
	ctx := context.Background()
	
	go func() {
		userMsg, _ := json.Marshal(schema.History{
			Role: "user",
			Content: userQuestion,
		})
		assistantMsg, _ := json.Marshal(schema.History{
			Role: "assistant",
			Content: answer.Answer,
		})
		pipe := redis.Rdb.Pipeline()
		pipe.RPush(ctx, key, userMsg, assistantMsg)
		pipe.LTrim(ctx, key, -10, -1)
		pipe.Expire(ctx, key, 20*time.Minute)
		pipe.Exec(ctx)
	}()

	// 消息持久化到数据库
	go message.CreateMessage(&schema.Message{
		SessionID: SessionID,
		UserName: session.UserName,
		Content: userQuestion,
		IsUser: true,
	})
	
	go message.CreateMessage(&schema.Message{
		SessionID: SessionID,
		UserName: session.UserName,
		Content: answer.Answer,
		IsUser: false,
	})

	return answer.Answer, code.CodeSuccess
}
