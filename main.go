package main

import (
	"urlShortener/controller"
	"urlShortener/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()

	r.POST("/add-url", controller.AddUrl)

	r.GET("/urls", controller.FindAllUrls)

	r.GET("/url/:urlId", controller.FindUrlByUrlId)

	r.Run()
}
