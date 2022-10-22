package service

import (
	"math/rand"
	"urlShortener/entity"
)

type Url struct {
	Urls map[string]string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (u *Url) Add(url entity.Url) entity.Url {
	u.Urls[RandStringBytes(6)] = url.Link

	return url
}

func (u *Url) FindAll() map[string]string {
	return u.Urls
}

func (u *Url) FindByShortenedId(id string) string {
	return u.Urls[id]
}
