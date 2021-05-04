package main

import "github.com/jinzhu/gorm"

func InitialMigration() {
	OpenDB()
	defer db.Close()
	db.AutoMigrate(&Book{}, &Author{})
}

func OpenDB() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Flied to connect to db")
	}
}
