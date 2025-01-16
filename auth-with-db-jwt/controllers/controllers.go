package controllers

import (
	"auth-with-db-jwt/authentication"
	"auth-with-db-jwt/database"
	"auth-with-db-jwt/models"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ServerUp(w http.ResponseWriter, r *http.Request) {
	// we will will write the header as Content-type application/json
	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"message": "Server is Live"})
}

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	// first we will set the headers
	w.Header().Set("Content-type", "application/json")

	// now we will read the data from the request
	requestData, err := io.ReadAll(r.Body)

	if checkError(w, err) {
		return
	}

	// now convert the request data to json
	var jsonData map[string]string

	json.Unmarshal(requestData, &jsonData)

	// now check if there is jwt token in the map or not
	_, ok := jsonData["jwtToken"]

	if !ok {
		checkError(w, errors.New("JWT Token not found"))
		return
	}

	// now we will verify the jwt token and return messages according to it
	jwtToken, err := authentication.VerifyToken(jsonData["jwtToken"])

	// if the token is invalid or expired it will throw an error
	if checkError(w, err) {
		return
	}

	username, _ := jwtToken.Claims.GetSubject()

	w.WriteHeader(http.StatusOK)
	var message string = "JWT token verified and valid"
	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
		"user":    username,
	})
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	// set the Content-type headers
	w.Header().Set("Content-type", "application/json")

	// User model
	var user models.User

	// make a new decoder and disallow all the unknown fields
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&user)

	if checkError(w, err) {
		return
	}

	result, err := database.SignUp(user)

	if checkError(w, err) {
		return
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	// after signup generate a jwt token for them
	jwtToken, err := authentication.GetToken(&user)
	if checkError(w, err) {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User signup successful", "jwtToken": jwtToken})
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	// first set the content headers
	w.Header().Set("Content-type", "application/json")

	var inputUser models.User

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&inputUser)

	if checkError(w, err) {
		return
	}

	// now we will get data from database
	databaseUser, err := database.SignIn(inputUser)
	if checkError(w, err) {
		return
	}

	// after we have got the user from the database
	// we will issue a jwt token to them

	jwtToken, err := authentication.GetToken(databaseUser)
	if checkError(w, err) {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "SignIn done", "jwtToken": jwtToken})
}

func checkError(w http.ResponseWriter, err error) bool {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "error " + err.Error()})

		return true
	}

	return false
}
