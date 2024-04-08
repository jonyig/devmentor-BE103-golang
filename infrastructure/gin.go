package infrastructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitGinServer() (server *gin.Engine, err error) {
	server = GinRouter()
	err = server.Run("127.0.0.1:8080")
	return
}

func GinRouter() (server *gin.Engine) {
	server = gin.New()
	server.GET("/test", test)

	return server
}

type TestData struct {
	Hello string `json:"hello"`
}

func test(c *gin.Context) {
	data := new(TestData)
	data.Hello = "world!"
	c.JSON(http.StatusOK, data)
}
