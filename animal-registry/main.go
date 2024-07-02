package main

import (
	"log"
	"net/http"

	"animal-registry/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/animal", handlers.RegisterAnimal).Methods("POST")
	router.HandleFunc("/api/animal/{id}", handlers.GetAnimal).Methods("GET")

	log.Println("API is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
