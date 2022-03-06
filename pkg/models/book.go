package models

import (
	"UnplugCharger/book-management/pkg/config"

	"gorm.io/gorm"
)


var db *gorm.DB



type Book struct{
	gorm.Model
	Name string `gorm:"json":"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`


}

func initDb()  {
	// Initialize db and automigrate
	config.Connect()
	db= config.GetDB()
	db.AutoMigrate(&Book{})
}