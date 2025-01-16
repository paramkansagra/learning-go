package main

import (
	"auth-with-db-jwt/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running main.go")

	fmt.Println("Server up at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", routers.Router()))
}
