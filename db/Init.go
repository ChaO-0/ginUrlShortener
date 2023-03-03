package db

import (
	"urlShortener/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	dsn := "host=pemendek-db user=postgres password=passw0rd dbname=goner port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Urls{})

	db = database
}

func GetDB() *gorm.DB {
	return db
}
