package kb

import (
	"github.com/google/uuid"
	"github.com/yuanji6666/gopherAI/common/mysql"
)

// Repository 数据访问层

// GetKnowledgeBaseByUsername 获取用户的所有知识库
func GetKnowledgeBaseByUsername(username string) ([]KnowledgeBase, error) {
	var kb []KnowledgeBase
	println(username)
	err := mysql.DB.Where("username = ?", username).Find(&kb).Error
	if err != nil {
		return nil, err
	}
	println(len(kb))
	return kb, nil
}

// CreateKnowledgeBase 创建新的知识库
func CreateKnowledgeBase(username string, kbName string) (UserKBID string, err error) {
	kb := KnowledgeBase{
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
