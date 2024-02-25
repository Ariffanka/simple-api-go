package routes

import (
	"api-go/controller"

	"github.com/gorilla/mux"

)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users/create", controller.Create).Methods("POST")
	router.HandleFunc("/users/{id}", controller.Get).Methods("GET")

	return router
}
