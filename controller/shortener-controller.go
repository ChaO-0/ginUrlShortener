package controller

import (
	"fmt"
	"net/http"
	"urlShortener/models"
	shortenerService "urlShortener/services"

	"github.com/gin-gonic/gin"
)

func AddUrl(c *gin.Context) {
	var url models.Urls

	fmt.Println(&url)

	c.BindJSON(&url)

	newUrl := shortenerService.Add(&url)

	c.JSON(http.StatusCreated, newUrl)
}

func FindAllUrls(c *gin.Context) {
	urls := shortenerService.FindAll()

	c.JSON(http.StatusOK, urls)
}

func FindUrlByUrlId(c *gin.Context) {
	urlId := c.Param("urlId")

	shortenedUrl := shortenerService.FindByShortenedId(urlId)

	gotoUrl := "https://gon.er/" + shortenedUrl.ShortenedId

	c.JSON(http.StatusOK, gin.H{
		"shortened_url": gotoUrl,
	})
}
