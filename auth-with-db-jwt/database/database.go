package database

import (
	"auth-with-db-jwt/models"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const databaseName string = "authWithDbJWT"
const collectionName string = "users"

var Collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(getConnectionString())

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Mongo client connection failed, err -> %v \n", err)
	}

	fmt.Println("Mongo Connection success")

	Collection = client.Database(databaseName).Collection(collectionName)

	fmt.Println("Connection reference Ready")
}

func getConnectionString() string {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("error loading .env file , error -> %v \n", err)
	}

	var mongodb_username string = os.Getenv("MONGODB_USERNAME")
	var mongodb_password string = os.Getenv("MONGODB_PASSWORD")
	var mongodb_url string = os.Getenv("MONGODB_URL")

	return fmt.Sprintf("mongodb+srv://%v:%v@%v", mongodb_username, mongodb_password, mongodb_url)
}

func checkSameUsernameEmailExists(username string, email string) error {
	usernameFilter := bson.D{{Key: "username", Value: username}}
	emailFilter := bson.D{{Key: "email", Value: email}}

	var user models.User

	err := Collection.FindOne(context.TODO(), usernameFilter).Decode(&user)
	if err == nil {
		return errors.New("account with same username found")
	}

	err = Collection.FindOne(context.TODO(), emailFilter).Decode(&user)
	if err == nil {
		return errors.New("account with same email found")
	}

	return nil
}

func checkPassword(inputUser *models.User, databaseUser *models.User) bool {
	// now we will hash the password from input and send to hash and check

	return models.CheckPassword(inputUser.Password, databaseUser.Password)
}

func SignUp(user models.User) (*mongo.InsertOneResult, error) {

	// first check if there are all the feilds or not
	err := models.CheckSignupRequiredFeilds(user)
	if err != nil {
		return nil, err
	}

	err = checkSameUsernameEmailExists(user.Username, user.Email)
	if err != nil {
		return nil, err
	}

	// first we will change the password to hashed password and then insert the data
	user.Password, err = models.HashUserPassword(user.Password)
	if err != nil {
		return nil, err
	}

	inserted, err := Collection.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	return inserted, nil
}

func SignIn(user models.User) (*models.User, error) {
	// we will either give that it was able to signin or not
	err := models.CheckSigninRequiredFeilds(user)
	if err != nil {
		return nil, err
	}

	// now we will try to sign them in using either the username or email
	usernameFilter := bson.D{{Key: "username", Value: user.Username}}
	emailFilter := bson.D{{Key: "email", Value: user.Email}}

	var databaseUser models.User

	if user.Username != "" {
		// get the result using email
		err := Collection.FindOne(context.TODO(), usernameFilter).Decode(&databaseUser)
		if err != nil {
			return nil, errors.New("account with that username not found")
		}
	} else {
		err := Collection.FindOne(context.TODO(), emailFilter).Decode(&databaseUser)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("account with that email not found")
		}
	}

	// now we will check for the id password
	if !checkPassword(&user, &databaseUser) {
		return nil, errors.New("wrong password")
	}

	return &databaseUser, nil
}
