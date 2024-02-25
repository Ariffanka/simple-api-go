package routes

import (
	"api-go/controller"

	"github.com/gorilla/mux"

)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users/create", controller.Create).Methods("POST")
	router.HandleFunc("/users/{id}", controller.Get).Methods("GET")
	router.HandleFunc("/users/", controller.Gets).Methods("GET")
	router.HandleFunc("/users/update/{id}", controller.Update).Methods("POST")
	router.HandleFunc("/users/delete/{id}", controller.Delete).Methods("GET")

	return router
}
