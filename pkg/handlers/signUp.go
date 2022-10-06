package handlers

import (
	"encoding/json"
	"example/restapi/pkg/models"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (h handler) SignUp(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	json.Unmarshal(body, &user)
	// Hashing password(it's workingworking)
	// HashedPassword, err := HashPassword(user.Password)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	log.Println("Internal Error Hash password")
	// }
	// user.Password = HashedPassword

	if result := h.DB.Create(&user); result.Error != nil {
		log.Fatal(result.Error)
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User Created")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
