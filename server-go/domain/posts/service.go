package posts

type PostService struct{
	pr *PostsRepo
}

func NewPostService(pr *PostsRepo) *PostService{
	return &PostService{pr: pr}
}

func (ps *PostService) PublishPost(username string, content string) error {
	post := &Post{
		Username: username, 
		Content: content,
		ViewCount: 0,
		LikeCount: 0,
		CommentCount: 0,
	}
	
	if err := ps.pr.PublishNewPost(post); err != nil {
		return nil
	}else{
		return err
	}
}