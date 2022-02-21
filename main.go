package main

import (
	// "encoding/json"
	"encoding/json"
	"log"
	"net/http"

	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)

// Book Struct
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// init books var as a slice Book struct
var books []Book

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	
}

// create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	
}

// update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	
}

// delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	// init router
	r := mux.NewRouter()

	// mock data
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "doe"}})
	books = append(books, Book{ID: "2", Isbn: "448744", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// route handlers / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8000", r))
}
