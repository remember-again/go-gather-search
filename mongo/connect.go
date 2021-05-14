package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client *mongo.Database

func newClient(dbname string) *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("connected mongodb error")
		return nil
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return nil
	}

	return client.Database(dbname)
}

func getDB(dbName string) *mongo.Database {
	if client == nil {
		client = newClient(dbName)
	}

	return client
}

func DB_FP() *mongo.Collection {
	return getDB("gather-search").Collection("favoritePage")
}
