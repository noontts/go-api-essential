package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "Java", Author: "Noon"})
	books = append(books, Book{ID: 2, Title: "Go", Author: "Noon"})

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("books/:id", deleteBook)

  app.Post("/upload", uploadFile)

	app.Listen(":8080")
}

func uploadFile(c *fiber.Ctx) error {
  file, err := c.FormFile("image")

  if err != nil {
    return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
  }

  err = c.SaveFile(file, "./uploads/" + file.Filename)

  if err != nil {
    return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
  }

  return c.SendString("File upload complete!")
}
