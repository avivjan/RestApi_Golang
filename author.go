package main

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
