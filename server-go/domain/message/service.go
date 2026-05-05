package message

import (
	"github.com/yuanji6666/gopherAI/common/code"
)

// Service 业务逻辑层

// GetHistoryBySessionIDWithID 根据会话ID和消息ID获取历史记录
func GetHistoryBySessionIDWithID(sessionID string, lastID int64, limit int) ([]History, code.Code) {
	msgs, err := GetMessagesBySessionID(sessionID, lastID, limit)

	if err != nil {
		return nil, code.CodeServerBusy
	}

	history := make([]History, len(msgs))

	for i, msg := range msgs {
		role := "assistant"
		if msg.IsUser {
			role = "user"
		}
		history[i] = History{
			ID:      msg.ID,
			Role:    role,
			Content: msg.Content,
		}
	}

	return history, code.CodeSuccess
}
