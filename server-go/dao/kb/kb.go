package kb

import (
	"github.com/google/uuid"
	"github.com/yuanji6666/gopherAI/common/mysql"
	"github.com/yuanji6666/gopherAI/schema"
)

func GetKnowledgeBaseByUsername(username string) ([]schema.KnowledgeBase, error) {
	var kb []schema.KnowledgeBase
	println(username)
	err := mysql.DB.Where("username = ?", username).Find(&kb).Error
	if err != nil {
		return nil, err
	}
	println(len(kb))
	return kb, nil
}

func CreateKnowledgeBase(username string, kbName string) (UserKBID string, err error) {
	kb := schema.KnowledgeBase{
		UserKBID: uuid.New().String(),
		Username: username,
		Name:     kbName,
	}
	err = mysql.DB.Create(&kb).Error
	if err != nil {
		return "", err
	}
	UserKBID = kb.UserKBID
	return UserKBID, nil
}
