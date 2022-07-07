package db

import (
	"context"
	"gradient-api/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToMongo() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)
	return client
}

var DB *mongo.Client = ConnectToMongo()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("gradient-dev").Collection(collectionName)
	return collection
}

var AlphaUsersCollection *mongo.Collection = GetCollection(DB, "alpha_users")
