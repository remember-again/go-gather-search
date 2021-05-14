package service

import (
	"zexho994/go-gather-search/entity"
	"zexho994/go-gather-search/form"
	"zexho994/go-gather-search/mongo"
)

func AddFpService(form form.AddFpForm) {

	entity := entity.FavoritePage{}
	entity.Score = form.Score
	entity.Email = form.Email
	entity.EmailType = form.EmailType
	entity.Url = form.Url
	entity.Title = form.Title

	mongo.AddFpToDB(entity, mongo.GetMongoDB("gather-search"))
}
