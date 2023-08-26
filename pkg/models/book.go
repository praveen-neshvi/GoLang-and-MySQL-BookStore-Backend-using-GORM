package models

import (
	"github.com/praveen-neshvi/bookstore/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books) //Books used here is a slice into which the results will be stored. GORM goes through the type of this slice which is Book (a model/struct defined by us) and hence it recognises that it must fetch data from books table of the database connected via db variable
	// var res []Book
	// res, _ := json.Marshal(Books)
	// fmt.Print(res)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) *Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return &book
}

// The line db.AutoMigrate(&Book{}) is used in GORM to automatically create or update
// database tables to match the schema defined by the specified GORM model. Here's what
// this line of code does:

// 1. db: This refers to the *gorm.DB instance, which represents a database connection managed by GORM.

// 2. AutoMigrate: This is a method provided by GORM that automatically creates or updates database tables based on the specified models.

// 3. &Book{}: This creates an instance of the Book struct. This instance is not actually used for anything other than providing GORM with the necessary information about the model's structure and fields.

// By calling db.AutoMigrate(&Book{}), you're instructing GORM to examine the Book model
// (struct) and compare it with the corresponding table in the database. GORM then makes
// the necessary changes to the database schema to ensure it matches the fields and
// structure of the Book model. This could involve creating the table if it doesn't
// exist or adding, modifying, or removing columns as needed.
