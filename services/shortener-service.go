package service

import (
	"math/rand"
	"urlShortener/db"
	"urlShortener/models"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Add(url *models.Urls) *models.Urls {
	dbInstance := db.GetDB()
	url.ShortenedId = RandStringBytes(6)

	dbInstance.Create(&url)

	return url
}

func FindAll() *[]models.Urls {
	var urls *[]models.Urls
	dbInstance := db.GetDB()

	dbInstance.Find(&urls)

	return urls
}

func FindByShortenedId(id string) *models.Urls {
	var url models.Urls
	dbInstance := db.GetDB()

	dbInstance.Where("shortened_id = ?", &id).First(&url)

	return &url
}
