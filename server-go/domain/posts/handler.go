package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/common/code"
)

type(

	PublishPostRequest struct{
		content string
	}
	
	PublishPostResponse struct{
		code.Response	
	}

)

type PostsHandler struct{
	ps *PostService
}

func NewPostsHandler(ps *PostService) *PostsHandler{
	return &PostsHandler{ps: ps}
}

func (ph *PostsHandler) PublishPost(c *gin.Context){

	req := new(PublishPostRequest)
	res := new(PublishPostResponse)

	username, exist := c.Get("username")
	if !exist{
		c.JSON(200, res.CodeOf(code.CodeServerBusy))
		return
	}
	
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	
	if err := ph.ps.PublishPost(username.(string), req.content); err != nil {
		c.JSON(200, res.CodeOf(code.CodeServerBusy))
		return 
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}