package service

import (
	"zexho994/go-gather-search/entity"
	"zexho994/go-gather-search/form"
	"zexho994/go-gather-search/mongo"
)

func AddFpService(form form.AddFpForm) {

	entity := entity.FavoritePage{}
	entity.Score = form.Score

	mongo.AddFpToDB()
}
