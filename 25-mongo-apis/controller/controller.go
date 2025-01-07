package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongo-api/models"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName string = "netflix"
const collectionName string = "watchlist"

// collection object for getting data from mongo db
var collection *mongo.Collection

// connect with mongoDB
func init() {
	// giving the client options
	clientOptions := options.Client().ApplyURI(getConnectionString())

	/*
		Context in hear means the different parameters like
			1. What is the deadline of the connection
			2. What are the different cancellation signals
			3. API boundries etc.

		Context.Background() -> this means the context is never cancelled and it keep on running in the background
		Context.TODO() -> this means the context is not clear ki muje kya kya use karna hai and kab tak karna hai
	*/

	// connect to mongo db
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Mongo client connection failed err -> %v \n", err)
	}

	fmt.Println("MongoDb connection success")

	collection = client.Database(dbName).Collection(collectionName)

	// collection reference is ready
	fmt.Println("Collection reference ready")
}

// making a getting a conection string using the env variables to change it independently
func getConnectionString() string {

	// first lets try to load the env file
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading the .env file")
	}

	var mongodb_username string = os.Getenv("MONGODB_USERNAME")
	var mongodb_password string = os.Getenv("MONGODB_PASSWORD")
	var mongodb_url string = os.Getenv("MONGODB_URL")
	return fmt.Sprintf("mongodb+srv://%v:%v@%v", mongodb_username, mongodb_password, mongodb_url)
}

// MongoDB helper functions -> seprate file usually
func insertOneMovie(movie models.Netflix) {
	/*
		so when we add the values we can have 2 things happening
			1. Successfully inserted the records
			2. Some kind of error that happened

		Case 1:-  If the data is successfully inserted then we will just fmt out the insertedID
		Case 2:- If we are not able to insert the data then we will log the error out
	*/
	inserted, err := collection.InsertOne(context.TODO(), movie)

	if err != nil {
		log.Fatalf("Not able to insert value err -> %v \n", err)
	}

	fmt.Println("Inserted one movie with id -> ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {

	/*
		so when we are trying to update the values there are many things to do
		1. Change the id from hex to primitive type that mongo accepts
			1.1 In case of any error during conversion from hex log it and mark it as fatal
		2. Then we have to set the filter i.e. which values to change
		3. Then we have to set what to update in values that we have found using filter
		4. Pass all of these values to mongoDb collection
	*/

	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatalf("Error found in converting movie id from hex -> %v \n", err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	updateResult, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatalf("Error recived during updating values -> %v \n", err)
	}

	fmt.Printf("Values updated -> %v \n", updateResult.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatalf("Error in converting movie id to hex -> %v \n", err)
	}

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatalf("Error in deleting the values -> %v \n", err)
	}

	fmt.Println("Number of movies that got deleted -> ", deleteResult.DeletedCount)
}

func deleteAllMovies() {
	filter := bson.D{}
	deletedResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error during deleting all movies -> %v \n", err)
	}

	fmt.Println("Number of deleted movies -> ", deletedResult.DeletedCount)
}

func getAllMovies() []models.Netflix {
	// get all the data from mongodb
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatalf("error in getting data from mongodb -> %v \n", err)
	}

	defer cursor.Close(context.Background()) // close the cursor after completing this function

	var movies []models.Netflix // making a array of primitive bson.M for our data which we will convert to objects in future

	// we are doing this on our own if any error comes please check the video and do it according to it
	for cursor.Next(context.Background()) {
		var movie models.Netflix
		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatalf("Error in decoding the bson -> %v \n", err)
		}

		movies = append(movies, movie)
	}

	return movies
}

// actual controllers :=

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Content-Allow-Methods", "GET")

	allMovies := getAllMovies()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Added the movie successfully")
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Movie watched updated successfully")
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Content-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Movie deleted successfully")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Content-Allow-Methods", "DELETE")

	deleteAllMovies()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("ALL movies deleted successfully")
}
