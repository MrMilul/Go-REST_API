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
func createBook(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {

	// Router
	r := mux.NewRouter()

	r.HandleFunc("/api/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	log.Println("Api is running")
	http.ListenAndServe(":8000", r)
}
