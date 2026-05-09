package posts

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Username 	 string `gorm:"varchar(50)" json:"username"`
	Content      string `gorm:"type:text" json:"content"`
	ViewCount    uint   `json:"view_count"`
	LikeCount    uint   `json:"like_count"`
	CommentCount uint   `json:"comment_count"`
}

type Comment struct {
	gorm.Model
	PostID    uint`gorm:"index" json:"post_id"`
	Username  string `gorm:"varchar(50)" json:"username"`
	Content   string `json:"content"`
	LikeCount uint   `json:"like_count"`
}
