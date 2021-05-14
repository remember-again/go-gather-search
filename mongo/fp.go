package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"zexho994/go-gather-search/entity"
	"zexho994/go-gather-search/util"
)

func AddFp(fp entity.FavoritePage) (res interface{}) {
	res, err := DB_FP().InsertOne(context.TODO(), fp)
	util.CheckErr(err, "insert one error")
	return
}

func FindFp(email string) []entity.FavoritePage {
	cur, err := DB_FP().Find(context.TODO(), bson.D{{"email", email}})
	util.CheckErr(err, "find fp error")
	var fps []entity.FavoritePage
	for cur.Next(context.TODO()) {
		var r entity.FavoritePage
		util.CheckErr(cur.Decode(&r), "decode fp error")
		fps = append(fps, r)
	}
	util.CheckErr(cur.Err(), "")
	util.CheckErr(cur.Close(context.TODO()), "")

	return fps
}
