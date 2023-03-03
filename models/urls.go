package models

import "gorm.io/gorm"

type Urls struct {
	gorm.Model
	Url         string `gorm:"type:text" json:"url"`
	ShortenedId string `gorm:"type:text" json:"shortened_url"`
}
