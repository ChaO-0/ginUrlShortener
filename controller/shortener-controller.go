package controller

import (
	"net/http"
	"urlShortener/entity"
	"urlShortener/services"

	"github.com/gin-gonic/gin"
)

var urlService service.Url = service.Url{}

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
