package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/mulugeta-89/library_rest_api_go_mysql/pkg/controllers"
)

var registerBookStoreRoutes = func (router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}

func main(){
	fmt.Println("Hello world")
}