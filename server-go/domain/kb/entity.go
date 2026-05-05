package kb

import "gorm.io/gorm"

// Entity 模型

// KnowledgeBase 知识库实体
type KnowledgeBase struct {
	gorm.Model
	UserKBID string `gorm:"type:varchar(36);uniqueIndex;not null" json:"user_kb_id"`
	Name     string `gorm:"type:varchar(255);not null;uniqueIndex:idx_username_name" json:"name"`
	Username string `gorm:"type:varchar(255);not null;uniqueIndex:idx_username_name" json:"username"`
}

// KnowledgeBaseInfo 知识库信息（用于API返回）
type KnowledgeBaseInfo struct {
	UserKBID string `json:"user_kb_id"`
	Name     string `json:"name"`
}
