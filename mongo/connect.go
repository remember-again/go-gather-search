package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Database

func newClient(dbname string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println("connected mongodb error")
		return nil
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}

	defer func(client *mongo.Client, ctx context.Context) {
		_ = client.Disconnect(ctx)
	}(client, ctx)

	db := client.Database(dbname)
	return db
}

func getClient(dbName string) *mongo.Database {
	if client == nil {
		client = newClient(dbName)
	}
	return client
}
