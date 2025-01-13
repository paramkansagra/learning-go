package auth_controllers

import (
	"encoding/json"
	"io"
	"jwt-tokens-with-apis/auth"
	"net/http"
)

func checkError(err error, w http.ResponseWriter) bool {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		error_info := map[string]string{"error": err.Error(), "message": "error occured"}
		json.NewEncoder(w).Encode(error_info)
		return true
	}

	return false
}

func CreateToken(w http.ResponseWriter, r *http.Request) {
	// we will set the headers of Content-type as application/json
	w.Header().Set("Content-type", "application/json")

	content, err := io.ReadAll(r.Body)

	if checkError(err, w) {
		return
	}

	var jsonData map[string]string
	err = json.Unmarshal(content, &jsonData)

	if checkError(err, w) {
		return
	}

	// now we will get the json from request

	var username string = jsonData["username"]
	var password string = jsonData["password"]

	jwtToken, err := auth.CreateToken(username, password)

	if checkError(err, w) {
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"jwtToken": jwtToken})
}

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	// first set the content type headers
	w.Header().Set("Content-type", "application/json")

	content, err := io.ReadAll(r.Body)

	if checkError(err, w) {
		return
	}

	// now we will check the request for some json data if possible
	var jsonData map[string]string
	err = json.Unmarshal(content, &jsonData)

	if checkError(err, w) {
		return
	}

	// now we will get the jwt token if present
	jwtToken, err := auth.VerifyToken(jsonData["jwtToken"])

	if checkError(err, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
	// write the jwtToken values to responce and send
	json.NewEncoder(w).Encode(jwtToken)
}
