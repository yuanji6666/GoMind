package posts

import (
	"gorm.io/gorm"
)

type PostsRepo struct{
	db *gorm.DB
}

func NewPostsRepo(db *gorm.DB) *PostsRepo{
	return &PostsRepo{db: db}
}

func (pr *PostsRepo) PublishNewPost(post *Post) error {
	return pr.db.Create(post).Error
}

func (pr *PostsRepo) GetPostsByUserID(userID uint) ([]Post, error){
	var posts []Post
	err := pr.db.Where("id = ?", userID).Find(&posts).Error
	return posts, err
}





