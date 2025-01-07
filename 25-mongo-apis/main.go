package main

import (
	"fmt"
	"log"
	"mongo-api/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")

	router := router.Router()

	log.Fatal(http.ListenAndServe(":4000", router))
	fmt.Println("Listning at port 4000")
}
