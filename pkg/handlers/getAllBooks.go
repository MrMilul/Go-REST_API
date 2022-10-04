package handlers

import (
	"encoding/json"
	"example/restapi/pkg/models"
	"fmt"
	"net/http"
)

func (h handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	var book []models.Book
	if result := h.DB.Find(&book); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
