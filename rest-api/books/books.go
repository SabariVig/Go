package books

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/sabarivig/go/rest-api/database"
)

//Book Struct 
type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Rating string `json:"rating"`
	Author string `json:"author"`
}

//GetBooks API ENDPOINT TO RETURN ALL BOOKS
func GetBooks(c *fiber.Ctx) {

	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

//GetBook API ENDPOINT TO RETURNS ONE BOOK
func GetBook(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("id")
	var book Book
	db.Find(&book, id)
	c.JSON(book)

}

//NewBook Creates A New Book
func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(404).Send(err)
		return
	}

	db.Create(&book)
	c.JSON(book)

}

//UpdateBook Updates The Book based on the book ID
func UpdateBook(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("id")
	book := new(Book)

	db.First(&book, id)
	if err := c.BodyParser(book); err != nil {
		c.Status(404).Send(err)
		return
	}
	db.Save(&book)
	c.JSON(book)

}

//DeleteBook Deletes A By Based On ID
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Name == "" {
		c.Status(404).Send("Book Not Found")
	}

	db.Delete(&book)
	c.Status(200).Send("Book Deleted Succesfully")
}
