package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
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

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {

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

	// route handlers / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8000", r))
}
