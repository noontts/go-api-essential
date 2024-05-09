package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Books struct {
  gorm.Model
	Name        string
	Author      string
	Description string
	Price       uint
}

func createBooks(db *gorm.DB, book *Books) {
  result := db.Create(book)

  if result.Error != nil {
    log.Fatalf("Error creating book: %v", result.Error)
  }

  fmt.Println("Create Book Successful")
}

func getBookss(db *gorm.DB, id uint) *Books{
  var book Books
  result := db.First(&book, id)

  if result.Error != nil {
    log.Fatalf("Error get book: %v", result.Error)
  }
  
  return &book
}

func updateBookss(db *gorm.DB,book *Books){
  result := db.Save(&book)

  if result.Error != nil {
    log.Fatalf("Error update book: %v", result.Error)
  }

  fmt.Println("Update Book Successful")
}
