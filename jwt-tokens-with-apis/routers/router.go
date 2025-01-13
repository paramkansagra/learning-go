package routers

import (
	auth_controllers "jwt-tokens-with-apis/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/createToken", auth_controllers.CreateToken).Methods("POST")
	r.HandleFunc("/verifyToken", auth_controllers.VerifyToken).Methods("POST")

	return r
}
