package controllers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func dbConnect() {

	DB_URL := "mongodb+srv://dexter:Password@mongogoapi@2@cluster0.d040v.mongodb.net/"
	var collection = *mongo.Collection

	// Set client options
	clientOptions := options.Client().ApplyURI(DB_URL)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection = client.Database("Netflix").Collection("Movies")

}
