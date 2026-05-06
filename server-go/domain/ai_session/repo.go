package ai_session

import (
	"slices"

	"github.com/yuanji6666/gopherAI/common/mysql"
)

// Repository 数据访问层

// CreateSession 创建新会话
func CreateSession(sessionEntity *Session) (*Session, error) {
	err := mysql.DB.Create(sessionEntity).Error
	return sessionEntity, err
}

// GetSessionByID 根据会话ID获取会话
func GetSessionByID(sessionID string) (Session, error) {
	var session Session
	err := mysql.DB.Where("id = ?", sessionID).First(&session).Error
	return session, err
}

// GetSessionByUsername 根据用户名获取会话列表
func GetSessionByUsername(username string) ([]Session, error) {
	var sessions []Session
	err := mysql.DB.Where("user_name = ?", username).Find(&sessions).Error
	return sessions, err
}

// CreateMessage 创建新消息
func CreateMessage(msg *Message) (*Message, error) {
	err := mysql.DB.Create(msg).Error
	return msg, err
}

// GetAllMessages 获取所有消息
func GetAllMessages() ([]Message, error) {
	var msgs []Message
	err := mysql.DB.Order("created_at asc").Find(&msgs).Error
	return msgs, err
}

// GetMessagesBySessionID 根据会话ID获取消息
func GetMessagesBySessionID(sessionID string, lastID int64, limit int) (msgs []Message, err error) {
	db := mysql.DB.Where("session_id = ?", sessionID)

	if lastID > 0 {
		db = db.Where("id < ?", lastID)
	}

	err = db.Order("id DESC").Limit(limit).Find(&msgs).Error
	if err != nil {
		return nil, err
	}

	slices.Reverse(msgs)

	return
}

// GetMessagesBySessionIDs 根据多个会话ID获取消息
func GetMessagesBySessionIDs(sessionIDs []string) (msgs []Message, err error) {
	err = mysql.DB.Where("session_id IN ?", sessionIDs).Order("created_at asc").Find(&msgs).Error
	return
}
