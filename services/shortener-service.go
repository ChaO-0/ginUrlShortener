package service

import "urlShortener/entity"

type Url struct {
	// urls []entity.Url
	urls map[string]string
}

var urls = make(map[string]string)

func (u *Url) Add(url entity.Url) entity.Url {
	// u.urls = append(u.urls, url)
	// u.urls = map[string]string{
	// 	url.Link: "willBeRandom",
	// }
	u.urls = urls{
		url.Link: "willBeRandom",
	}

	return url
}

func (u *Url) FindAll() map[string]string {
	return u.urls
}
