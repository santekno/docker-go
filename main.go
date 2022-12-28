package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/santekno/golang-app/internal/handler"
)

func main() {
	fmt.Println("GO Docker Tutorial CI/CD")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	r.GET("/world", handler.HelloWorld)
	r.POST("/sum", handler.SumHandler)
	r.POST("/multiply", handler.MultiplicationHandler)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
