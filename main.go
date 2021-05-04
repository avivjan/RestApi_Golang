package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {
	//Init Router
	router := mux.NewRouter()

	InitialMigration()

	//Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func InitialMigration() {
	db, err = gorm.Open("sqlite3", "test.db")
	OpenDB()
	defer db.Close()
	db.AutoMigrate(&Book{}, &Author{})
}
