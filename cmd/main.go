package main

import (
	"example/restapi/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getBook(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {

	// Router
	r := mux.NewRouter()

	// Handlers
	r.HandleFunc("/api/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/api/books", handlers.CreateBook).Methods("POST")

	log.Println("Api is running")
	http.ListenAndServe(":8000", r)
}
