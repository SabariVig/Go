package main

import (

	// "fmt"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sabarivig/go/rest-api/books"
	"github.com/sabarivig/go/rest-api/database"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello!!")

}

func initilizeDB() {
	var err error
	database.DBConn, err = gorm.Open("postgres", os.Getenv("POSTGRESS_URI"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&books.Book{})
	fmt.Println("Migration Completed")
}

func hello(c *fiber.Ctx) {

	c.Send("Hello")

}
func routes(c *fiber.App) {
	c.Get("/books", books.GetBooks)
	c.Get("/", hello)
	c.Get("/book/:id", books.GetBook)
	c.Put("/book/", books.NewBook)
	c.Put("/book/:id", books.UpdateBook)
	c.Delete("/book/:id", books.DeleteBook)
}

func main() {
	initilizeDB()
	app := fiber.New()
	routes(app)
	defer database.DBConn.Close()
	app.Listen(3000)
}
