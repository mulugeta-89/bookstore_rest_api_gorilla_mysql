package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mulugeta-89/library_rest_api_go_mysql/pkg/config"
)

var db *gorm.DB
type Book struct {
	gorm.Model
	Name string `gorm: ""json: "name"`
	Author string `json: "author"`
	Publication string `json: "publication"`
}

func init () {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}

func DeleteBooks() {
	db.Delete(&Book{})
}

