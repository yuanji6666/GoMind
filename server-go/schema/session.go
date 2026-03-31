package schema

import (
	"time"

	"gorm.io/gorm"
)

// Session UserKBID 每个回话对应唯一数据库id
type Session struct {
	ID        string         `gorm:"primaryKey;type:varchar(36)" json:"id"`
	UserName  string         `gorm:"index;not null" json:"username"`
	UserKBID  string         `gorm:"type:varchar(36);not null" json:"user_kb_id"`
	Title     string         `gorm:"type:varchar(100)" json:"title"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type SessionInfo struct {
	Title string `json:"title"`
	SessionID string `json:"session_id"`
	UserKBID string `json:"user_kb_id"`
}