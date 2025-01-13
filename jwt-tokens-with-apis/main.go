package main

import (
	"jwt-tokens-with-apis/routers"
	"log"
	"net/http"
)

func main() {

	router := routers.Router()

	log.Fatal(http.ListenAndServe(":8000", router))
}
