package handlers

import (
	"encoding/json"
	"example/restapi/pkg/mocks"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mocks.Books)
}
