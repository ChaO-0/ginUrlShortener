package main

import (
	"urlShortener/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/add-url", controller.AddUrl)

	r.GET("/urls", controller.FindAllUrls)

	r.GET("/url/:urlId", controller.FindUrlByUrlId)

	r.Run()
}
