package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {

	client, err := *mongo.NewClient(options.Client().ApplyURl("mongodb://localhost:3000"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Prinln("Failed to Connect to Database :(")
		return nil
	}
	fmt.Println("Successfully Connected to Database")
	return client
}

var client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionname string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionname)
	return collection

}

func ProductData(client *mongo.Client, collectionname string) *mongo.Collection {

	var productCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionname)
	return productCollection
}
