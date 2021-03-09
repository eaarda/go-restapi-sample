package Models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Page   int    `json:"page"`
}
