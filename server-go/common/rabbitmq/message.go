// message.go 定义消息格式，支持消息序列化，定义消息处理方法
package rabbitmq

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yuanji6666/gopherAI/dao/message"
	"github.com/yuanji6666/gopherAI/schema"
)

// MessageMQParam 消息格式
type MessageMQParam struct {
	SessionID string `json:"session_id"`
	Content   string `json:"content"`
	UserName  string `json:"user_name"`
	IsUser    bool   `json:"is_user"`
}

// GenerateMessageMQParam 消息序列化
func GenerateMessageMQParam(sessionID, content, username string, isUser bool) []byte {
	messageMQParam := &MessageMQParam{
		SessionID: sessionID,
		Content:   content,
		UserName:  username,
		IsUser:    isUser,
	}

	bytes, _ := json.Marshal(messageMQParam)
	return bytes
}

// HandleMQMessage 消费者侧消息处理
func HandleMQMessage(msg *amqp.Delivery) error {
	var param MessageMQParam
	err := json.Unmarshal(msg.Body, &param)

	if err != nil {
		return err
	}

	newMsg := &schema.Message{
		SessionID: param.SessionID,
		Content:   param.Content,
		UserName:  param.UserName,
		IsUser:    param.IsUser,
	}

	_, err = message.CreateMessage(newMsg)

	return err
}
