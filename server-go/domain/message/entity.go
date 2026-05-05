package message

import (
	"time"
)

// Entity 模型

// Message 消息实体
type Message struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	SessionID string    `gorm:"index;not null;type:varchar(36)" json:"session_id"`
	UserName  string    `gorm:"type:varchar(20)" json:"username"`
	Content   string    `gorm:"type:text" json:"content"`
	IsUser    bool      `gorm:"not null;" json:"is_user"`
	CreatedAt time.Time `json:"created_at"`
}

// History 历史记录（用于API返回）
type History struct {
	ID      uint   `json:"id"`
	Role    string `json:"role"`
	Content string `json:"content"`
}
