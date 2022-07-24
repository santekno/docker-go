package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorld is sample API Hello world
func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello World",
	})
}
