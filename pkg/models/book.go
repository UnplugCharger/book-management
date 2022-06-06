package models

import (
	config "github.com/UnplugCharger/book-management/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB


type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	// Initialize db and automigrate
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//struct method to create a book entry inside the database
func (b *Book) CreateBook() *Book {

	db.Create(&b)
	return b
}

// GetAllBooks fetch a slice of all books from the database
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
