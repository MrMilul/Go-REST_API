package handlers

import (
	"encoding/json"
	"example/restapi/pkg/models"
	"example/restapi/pkg/utils"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret-key")

type Cridentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	var user Cridentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}
	var User models.User
	if result := h.DB.Where("Username=?", user.Username).Find(&User); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("User does not exist!"))
		return
	}

	if !utils.CheckHashedPassword(user.Password, User.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid Cridentials"))
		return
	}

	expireTime := time.Now().Add(time.Minute * 5)

	claims := Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   tokenString,
		Expires: expireTime,
	})

}
