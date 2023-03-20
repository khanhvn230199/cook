package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("mysql", "root:123456789@tcp(127.0.0.1:3306)/cook")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Cook{})

	DB = database
}
