package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/common/mysql"
	"github.com/yuanji6666/gopherAI/domain/posts"
)
var db = mysql.DB
var pr = posts.NewPostsRepo(db) 
var ps = posts.NewPostService(pr)
var ph = posts.NewPostsHandler(ps)

func RegisterPostRouter(r *gin.RouterGroup){
	r.POST("/post", ph.PublishPost)
}