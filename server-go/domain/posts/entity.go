package posts

import (
	"gorm.io/gorm"
)

type Post struct{
	gorm.Model
	PostID int `gorm:"index"` 
	UserID int 
	Content string `gorm:"type:text"`
	ViewCount int
	LikeCount int 
	CommentCount int
}

type Comment struct{
	gorm.Model
	
}

