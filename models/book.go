package Models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Page   int    `json:"page"`
}
