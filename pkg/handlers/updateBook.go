package handlers

import (
	"encoding/json"
	"example/restapi/pkg/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	var book models.Book
	if result := h.DB.Find(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Desc = updatedBook.Desc
	h.DB.Save(&book)

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("updated")
}
