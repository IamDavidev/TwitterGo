package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection = ConnectionToMongoDb()

var clientOptions = options.Client().ApplyURI("mongodb+srv://root:rootPass@cluster0.oklsp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func ConnectionToMongoDb() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("connected to mongoDB", client)
	return client

}

func IsConnection() bool {

	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	return true

}
