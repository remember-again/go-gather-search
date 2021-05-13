package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"zexho994/go-gather-search/entity"
	"zexho994/go-gather-search/util"
)

func AddFpToDB(fp entity.FavoritePage, db *mongo.Database) (res interface{}) {
	res, err := db.Collection("favoritePage").InsertOne(context.TODO(), fp)
	util.CheckErr(err, "insert one error")
	return
}
