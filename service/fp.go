package service

import (
	"zexho994/go-gather-search/entity"
	"zexho994/go-gather-search/form"
	mongo "zexho994/go-gather-search/mongo"
)

func AddFp(form form.AddFpForm) {
	e := entity.FavoritePage{}
	e.Score = form.Score
	e.Email = form.Email
	e.EmailType = form.EmailType
	e.Url = form.Url
	e.Title = form.Title

	mongo.AddFp(e)
}

func FindFp(email string) []entity.FavoritePage {
	return mongo.FindFp(email)
}

func RemoveFp(email, url string) bool {
	return mongo.RemoveFp(email, url)
}
