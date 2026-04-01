package message

import (
	"slices"

	"github.com/yuanji6666/gopherAI/common/mysql"
	"github.com/yuanji6666/gopherAI/schema"
)

func CreateMessage(msg *schema.Message) (*schema.Message, error) {
	err := mysql.DB.Create(msg).Error
	return msg, err
}

func GetAllMessages() ([]schema.Message, error) {
	var msgs []schema.Message
	err := mysql.DB.Order("created_at asc").Find(&msgs).Error
	return msgs, err
}

func GetMessagesBySessionID(sessionID string, lastID int64, limit int) (msgs []schema.Message, err error) {
	db := mysql.DB.Where("session_id = ?", sessionID)
	
	if lastID > 0 {
		db = db.Where("id < ?", lastID)
	}
	
	err = db.Order("id DESC").Limit(limit).Find(&msgs).Error
	if err != nil{
		return nil, err
	}

	slices.Reverse(msgs)

	return
}

func GetMessagesBySessionIDs(sessionIDs []string) (msgs []schema.Message, err error) {
	err = mysql.DB.Where("session_id IN ?", sessionIDs).Order("created_at asc").Find(&msgs).Error
	return
}

