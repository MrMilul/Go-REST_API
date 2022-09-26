package handlers

import (
	"encoding/json"
	"example/restapi/pkg/mocks"
	"example/restapi/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var book models.Books
	json.Unmarshal(body, &book)
	book.ID = "2"
	mocks.Books = append(mocks.Books, book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("A book created")

}
