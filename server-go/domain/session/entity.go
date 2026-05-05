package session

import (
	"time"

	"gorm.io/gorm"
)

// Entity 模型

// Session 会话实体 - UserKBID 每个会话对应唯一数据库id
type Session struct {
	ID        string         `gorm:"primaryKey;type:varchar(36)" json:"id"`
	UserName  string         `gorm:"index;not null" json:"username"`
	UserKBID  string         `gorm:"type:varchar(36);not null" json:"user_kb_id"`
	Title     string         `gorm:"type:varchar(100)" json:"title"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// SessionInfo 会话信息（用于API返回）
type SessionInfo struct {
	Title     string `json:"title"`
	SessionID string `json:"session_id"`
	UserKBID  string `json:"user_kb_id"`
}
