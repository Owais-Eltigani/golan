package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongo/modules/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"
)

//* DB helper methods

var collection *mongo.Collection
var DB_URL = "mongodb+srv://dexter:Password%40mongogoapi%402@cluster0.d040v.mongodb.net/"
var DB_NAME = "Netflix"
var COLLECTION_NAME = "Movies"
var client *mongo.Client

func DB_Connect() {

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

	collection = client.Database(DB_NAME).Collection(COLLECTION_NAME)
	fmt.Println("Connected to MongoDB!")

}

func DB_Disconnect() error {

	// already there is a connection
	if client != nil {
		// Using context with timeout for clean shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := client.Disconnect(ctx); err != nil {
			return fmt.Errorf("failed to disconnect from MongoDB: %v", err)
		}
		fmt.Println("Successfully disconnected from MongoDB!")
		return nil
	}
	return nil
}

func insertOneMovie(movie models.Netflix) {

	//
	inserted, err := collection.InsertOne(context.TODO(), movie) //my test code.
	// collection.InsertOne(context.Background(), movie) //correct code

	if err != nil {
		log.Fatal("couldn't insert movie.\n")
		return
	}

	fmt.Printf("movie inserted with id: %v\n", inserted.InsertedID)
}

func updateOneMove(movieID string) {

	id, err := primitive.ObjectIDFromHex(movieID)

	if err != nil {
		log.Fatal("couldn't convert movieId: ", movieID, "\nparsed_id: ", id, "\n")

	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result := collection.FindOneAndUpdate(context.Background(), filter, update)

	if result == nil {

		log.Fatal("couldn't update movie watched attribute.\n", result)
		return
	}

	fmt.Printf("movie watched updated.\n")

}

func deleteOneMovie(movieID string) {

	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	result, _ := collection.DeleteOne(context.Background(), filter)

	if result == nil {

		log.Fatal("couldn't delete watched information.\n")
		return
	}

	fmt.Printf("watched updated.\n")

}

func deleteAllMovies() int64 {

	filter := bson.M{}
	deletedMovies, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal("couldn't delete all movies,\n")
	}

	fmt.Printf("all movies have been deleted, %v in total.\n", deletedMovies.DeletedCount)
	return deletedMovies.DeletedCount
}

func getAllMovies() ([]primitive.M, error) {

	cursor, err := collection.Find(context.Background(), bson.M{})
	var movies []primitive.M

	if err != nil {
		log.Fatal("something went wrong while fetching movies.\n", err, movies)
		return movies, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {

		var movie primitive.M
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)

	}
	return movies, err

}

//* DB helper methods

// * actual controllers start here

// GetAllDbMovies : fetch all movies in db
func GetAllDbMovies(wtr http.ResponseWriter, req *http.Request) {

	allMovies, err := getAllMovies()

	if err != nil {
		fmt.Printf("something went wrong while trying fetching data")
		json.NewEncoder(wtr).Encode("something went wrong while trying fetching data")
		return
	}
	wtr.Header().Set("Content-Type", "application/x-www-form-urlencode")

	json.NewEncoder(wtr).Encode(allMovies)

}

// add a new movie to db
func InsertOneMovie(wtr http.ResponseWriter, req *http.Request) {
	wtr.Header().Set("Content-Type", "application/json")

	var movie models.Netflix
	json.NewDecoder(req.Body).Decode(&movie)

	insertOneMovie(movie)
	json.NewEncoder(wtr).Encode(movie)
}

func MarkOneMovieWatched(wtr http.ResponseWriter, req *http.Request) {
	wtr.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	updateOneMove(params["id"])

	json.NewEncoder(wtr).Encode("updated.\n")
}

func DeleteOneMovie(wtr http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	id := params["id"]

	deleteOneMovie(id)
	json.NewEncoder(wtr).Encode("deleted.\n")
}

func DeleteAll(wtr http.ResponseWriter, req *http.Request) {

	count := deleteAllMovies()
	json.NewEncoder(wtr).Encode(count)
}

//* actual controllers start here
