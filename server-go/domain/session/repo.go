package session

import (
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
