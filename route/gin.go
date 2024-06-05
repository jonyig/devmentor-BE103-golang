package route

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/handler/general"
	"shopping-cart/handler/post"
	"shopping-cart/handler/user"
)

func InitGinServer() (server *gin.Engine, err error) {
	server = GinRouter()
	err = server.Run("127.0.0.1:8080")
	return
}

func GinRouter() (server *gin.Engine) {
	server = gin.New()

	server.LoadHTMLGlob("frontend/*")

	api := server.Group("/api")
	//orderGroup := api.Group("/orders")
	//userGroup := api.Group("/auth")

	group := server.Group("")

	post.NewPosts(group)
	general.NewGeneral(group)
	user.NewAuthorization(api)

	return server
}
