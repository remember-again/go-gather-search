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

func RemoveFp(email, url string) bool {
	filter := bson.D{{"email", email}, {"url", url}}
	_, err := DB_FP().DeleteOne(context.TODO(), filter)
	if err != nil {
		return false
	}
	return true
}

func FindOne(email, url string) *entity.FavoritePage {
	filter := bson.D{{"email", email}, {"url", url}}
	//fp = entity.FavoritePage{}
	fp := &entity.FavoritePage{}
	err := DB_FP().FindOne(context.TODO(), filter).Decode(fp)
	if err != nil {
		return nil
	}
	return fp
}
