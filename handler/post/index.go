package post

import (
	"devmentor-BE103-golang/repository"
	"github.com/gin-gonic/gin"
)

type Post struct {
	postRepository repository.PostRepositoryInterface
}

func NewPosts(
	r *gin.RouterGroup,

) *Post {
	h := &Post{
		postRepository: repository.NewPostRepository(),
	}

	newRoute(h, r)

	return h
}

func newRoute(h *Post, r *gin.RouterGroup) {
	Group := r.Group("posts")

	Group.GET("", h.get)
	Group.POST("", h.create)
}
