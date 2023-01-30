package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type numbersBody struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func SumHandler(c *gin.Context) {
	var body numbersBody
	if err := c.ShouldBindJSON(&body); err != nil {
		err = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": body.X + body.Y,
	})
}

func MultiplicationHandler(c *gin.Context) {
	var body numbersBody
	if err := c.ShouldBindJSON(&body); err != nil {
		err = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": body.X * body.Y,
	})
}
