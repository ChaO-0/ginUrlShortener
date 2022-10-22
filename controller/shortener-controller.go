package controller

import (
	"net/http"
	"urlShortener/entity"
	"urlShortener/services"

	"github.com/gin-gonic/gin"
)

var urlService service.Url = service.Url{
	Urls: make(map[string]string),
}

func AddUrl(c *gin.Context) {
	var url entity.Url
	c.BindJSON(&url)

	newUrl := urlService.Add(url)

	c.JSON(http.StatusCreated, newUrl)
}

func FindAllUrls(c *gin.Context) {
	urls := urlService.FindAll()

	c.JSON(http.StatusOK, urls)
}

func FindUrlByUrlId(c *gin.Context) {
	urlId := c.Param("urlId")

	goToUrl := urlService.FindByShortenedId(urlId)

	c.JSON(http.StatusOK, gin.H{
		"goto": goToUrl,
	})
}
