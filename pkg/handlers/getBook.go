package handlers

import (
	"encoding/json"
	"example/restapi/pkg/mocks"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range mocks.Books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}

	}

}
