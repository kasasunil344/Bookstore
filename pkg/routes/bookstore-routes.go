package routes

import (
	"github.com/gorilla/mux"
	"github.com/kasasunil344/go-bookstore/pkg/controllers"
)

func RegisterBookStoreRoutes(r *mux.Router) {
	r.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
