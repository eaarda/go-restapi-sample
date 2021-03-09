package Config

import (
	"api-test/Models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB
var err error

func InitialMigration() {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to db")
	}

	database.AutoMigrate(&Models.Book{})
	DB = database
}
