package models

import (
	"github.com/Sparsh1401/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var Db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	Db = config.GetDB()

	Db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	Db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	Db.Find(&Books)
	return Books
}

func GetBookByID(Id int64) (*Book, *gorm.DB) {
	var book Book
	Db = Db.Where("Id=?", Id).Find(&book)
	return &book, Db
}

func DeleteBook(Id int64) *Book {
	var book Book
	Db.Where("Id=?", Id).Delete(book)
	return &book
}
