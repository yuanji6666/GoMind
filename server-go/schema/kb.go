package schema

import "gorm.io/gorm"

type KnowledgeBase struct {
	gorm.Model
	UserKBID string `gorm:"type:varchar(36);uniqueIndex;not null" json:"user_kb_id"`
	Name     string `gorm:"type:varchar(255);not null;uniqueIndex:idx_username_name" json:"name"`
	Username string `gorm:"type:varchar(255);not null;uniqueIndex:idx_username_name" json:"username"`
}

type KnowledgeBaseInfo struct {
	UserKBID string `json:"user_kb_id"`
	Name     string `json:"name"`
}
