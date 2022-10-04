package main

import (
	"example/restapi/pkg/db"
	"example/restapi/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// DB initialization
	DB := db.Init()
	h := handlers.New(DB)
	// Router
	r := mux.NewRouter()
	// Handlers
	r.HandleFunc("/api/books", h.GetBooks).Methods(http.MethodGet)
	r.HandleFunc("/api/book/{id}", h.GetBook).Methods(http.MethodGet)
	r.HandleFunc("/api/books", h.CreateBook).Methods(http.MethodPost)
	r.HandleFunc("api/book/{id}", h.DeleteBook).Methods(http.MethodDelete)
	r.HandleFunc("api/book/{id}", h.UpdateBook).Methods(http.MethodPut)

	log.Println("Api is running")
	http.ListenAndServe(":8000", r)
}
