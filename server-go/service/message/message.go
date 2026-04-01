package message

import (
	"github.com/yuanji6666/gopherAI/common/code"
	"github.com/yuanji6666/gopherAI/dao/message"
	"github.com/yuanji6666/gopherAI/schema"
)

func GetHistoryBySessionIDWithID(sessionID string, lastID int64, limit int) ([]schema.History, code.Code){
	msgs, err := message.GetMessagesBySessionID(sessionID, lastID, limit)
	
	if err != nil {
		return nil, code.CodeServerBusy
	}
	
	history := make([]schema.History, len(msgs))

	for i, msg := range msgs {
		role := "assistant"
		if msg.IsUser {
			role = "user"
		}
		history[i] = schema.History{
			ID:      msg.ID,
			Role:    role,
			Content: msg.Content,
		}
	}
	
	return history, code.CodeSuccess
}