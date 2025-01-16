package routers

import (
	"auth-with-db-jwt/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func Router() *mux.Router {
	router := mux.NewRouter()

	router.Use(loggingMiddleware)

	router.HandleFunc("/", controllers.ServerUp).Methods("GET")
	router.HandleFunc("/verifyToken", controllers.VerifyToken).Methods("POST")
	router.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/signin", controllers.SignIn).Methods("POST")

	return router
}
