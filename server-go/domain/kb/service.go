package kb

import (
	"github.com/yuanji6666/gopherAI/common/code"
)

// Service 业务逻辑层

// GetKnowledgeBaseInfoList 获取用户知识库列表
func GetKnowledgeBaseInfoList(username string) ([]KnowledgeBaseInfo, code.Code) {
	kbList, err := GetKnowledgeBaseByUsername(username)
	if err != nil {
		return nil, code.CodeServerBusy
	}

	var infoList []KnowledgeBaseInfo
	for _, kb := range kbList {
		infoList = append(infoList, KnowledgeBaseInfo{
			UserKBID: kb.UserKBID,
			Name:     kb.Name,
		})
	}

	return infoList, code.CodeSuccess
}

// CreateKB 创建知识库
func CreateKB(username string, kbName string) (KnowledgeBaseInfo, code.Code) {
	userKBID, err := CreateKnowledgeBase(username, kbName)
	if err != nil {
		return KnowledgeBaseInfo{}, code.CodeServerBusy
	}
	return KnowledgeBaseInfo{
		UserKBID: userKBID,
		Name:     kbName,
	}, code.CodeSuccess
}
