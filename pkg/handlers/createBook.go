package handlers

import (
	"encoding/json"
	"example/restapi/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) CreateBook(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)
	// book.ID = strconv.Itoa(rand.Intn(10000))
	// mocks.Books = append(mocks.Books, book)
	if result := h.DB.Create(&book); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(result.Error)
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("A book created")

}
