package session

import (
	"github.com/yuanji6666/gopherAI/common/mysql"
	"github.com/yuanji6666/gopherAI/schema"
)

func CreateSession(session *schema.Session) (*schema.Session, error) {
	err := mysql.DB.Create(session).Error
	return session, err
}

func GetSessionByID(sessionID string) (schema.Session, error) {
	var session schema.Session
	err := mysql.DB.Where("id = ?", sessionID).First(&session).Error
	return session, err
}

func GetSessionByUsername(username string) ([]schema.Session, error) {
	var sessions []schema.Session
	err := mysql.DB.Where("user_name = ?", username).Find(&sessions).Error
	return sessions, err
}
