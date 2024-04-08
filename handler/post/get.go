package post

import (
	"devmentor-BE103-golang/model/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Post) get(c *gin.Context) {
	posts := database.Posts{}

	err := posts.FindAll()
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}
