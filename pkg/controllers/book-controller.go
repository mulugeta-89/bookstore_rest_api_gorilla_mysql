package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mulugeta-89/library_rest_api_go_mysql/pkg/models"
	"github.com/mulugeta-89/library_rest_api_go_mysql/pkg/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	allBooks := models.GetBooks();
	res, _ := json.Marshal(allBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	models.DeleteBooks()
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "[]")
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	bookId := Vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book, _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createbook := &models.Book{}
	utils.ParseBody(r, createbook)
	b := createbook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	bookId := Vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing the data")
	}

	deleted := models.DeleteBook(ID)
	res, _ := json.Marshal(deleted)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatedbook := &models.Book{}
	utils.ParseBody(r, updatedbook)

	Vars := mux.Vars(r)
	bookId := Vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)

	if err != nil {
		fmt.Println("error while parsing")
	}
	theBook, db := models.GetBookById(ID)

	if updatedbook.Name != "" {
		theBook.Name = updatedbook.Name
	}
	if updatedbook.Author != "" {
		theBook.Author = updatedbook.Author
	}
	if updatedbook.Publication != "" {
		theBook.Publication = updatedbook.Publication
	}

	db.Save(&theBook)

	res, _ := json.Marshal(theBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}