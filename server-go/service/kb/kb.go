package kb

import (
	"github.com/yuanji6666/gopherAI/common/code"
	"github.com/yuanji6666/gopherAI/dao/kb"
	"github.com/yuanji6666/gopherAI/schema"
)

func GetKnowledgeBaseInfoList(username string) ([]schema.KnowledgeBaseInfo, code.Code) {
	kbList, err := kb.GetKnowledgeBaseByUsername(username)
	if err != nil {
		return nil, code.CodeServerBusy
	}

	var infoList []schema.KnowledgeBaseInfo
	for _, kb := range kbList {
		infoList = append(infoList, schema.KnowledgeBaseInfo{
			UserKBID: kb.UserKBID,
			Name: kb.Name,
		})
	}

	return infoList, code.CodeSuccess
}

func CreateKnowledgeBase(username string, kbName string) (schema.KnowledgeBaseInfo, code.Code) {
	userKBID, err := kb.CreateKnowledgeBase(username, kbName)
	if err != nil {
		return schema.KnowledgeBaseInfo{}, code.CodeServerBusy
	}
	return schema.KnowledgeBaseInfo{
		UserKBID: userKBID,
		Name: kbName,
	}, code.CodeSuccess
}