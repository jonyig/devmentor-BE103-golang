package post

import "github.com/gin-gonic/gin"

type Post struct {
}

func NewPosts(
	r *gin.RouterGroup,

) *Post {
	h := &Post{}

	newRoute(h, r)

	return h
}

func newRoute(h *Post, r *gin.RouterGroup) {
	Group := r.Group("posts")

	Group.GET("", h.get)
	Group.POST("", h.create)
}
