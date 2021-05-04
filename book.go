package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Book struct {
	gorm.Model
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	OpenDB()
	defer db.Debug().Close()
	var books []Book
	db.Debug().Find(&books)
	json.NewEncoder(w).Encode(books)

}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	requestedId := params["id"]
	OpenDB()
	defer db.Debug().Close()
	var book Book
	if !db.Debug().First(&book, requestedId).RecordNotFound() {
		json.NewEncoder(w).Encode(book)
		return
	}
	json.NewEncoder(w).Encode("There is no book with this id")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	OpenDB()
	defer db.Debug().Close()
	db.Debug().Create(&book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	requestedId := params["id"]
	var updatedBook Book
	OpenDB()
	defer db.Debug().Close()
	if db.Debug().First(&updatedBook, requestedId).RecordNotFound() {
		json.NewEncoder(w).Encode("There is no book with this id")
		return
	}
	_ = json.NewDecoder(r.Body).Decode(&updatedBook)
	db.Debug().Save(updatedBook)
	json.NewEncoder(w).Encode(updatedBook)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	requestedId := params["id"]
	var book Book
	OpenDB()
	defer db.Debug().Close()
	db.Debug().First(&book, requestedId)
	db.Debug().Delete(book)
	json.NewEncoder(w).Encode("Book deleted!")

}
