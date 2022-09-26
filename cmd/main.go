package main

import (
	"encoding/json"
	"example/restapi/pkg/mocks"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mocks.Books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	return
}
func createBook(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {

	// Router
	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	log.Println("Api is running")
	http.ListenAndServe(":8000", r)
}
