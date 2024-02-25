package main

import (
	"api-go/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	log.Println("server running on port 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
