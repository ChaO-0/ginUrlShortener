package main

import (
	"net/http"
	"urlShortener/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/shorten", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/add-url", controller.AddUrl)

	r.GET("/urls", controller.FindAllUrls)

	r.Run()
}
