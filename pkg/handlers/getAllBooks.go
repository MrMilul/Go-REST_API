package handlers

import (
	"encoding/json"
	"example/restapi/pkg/models"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func (h handler) GetBooks(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("Token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Please login first"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("First bad request"))
		return
	}

	tokenStr := cookie.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("second bad request"))
		log.Println(err)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var book []models.Book
	if result := h.DB.Find(&book); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
