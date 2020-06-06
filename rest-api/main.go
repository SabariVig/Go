package main

import (
	"log"
	// "fmt"
	"fmt"

	"github.com/go-bongo/bongo"
	"github.com/gofiber/fiber"
	"github.com/sabarivig/go/rest-api/books"
	"github.com/sabarivig/go/rest-api/database"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type user struct {
	bongo.DocumentBase `bson:",inline"`
	name               string
	pass               string
}

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello!!")

}

func routes(c *fiber.App) {
	c.Get("/books", books.GetBooks)
	c.Get("/book/:id", books.GetBook)
	c.Post("/book/:id", books.NewBook)
	c.Delete("/book/:id", books.DeleteBook)
}
func databaseTese() {
	fmt.Print("Hello")
	err := database.Client.Connect(database.Ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Client.Disconnect(database.Ctx)
	err = database.Client.Ping(database.Ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	databaseTese()
	app := fiber.New()
	routes(app)
	app.Listen(3000)
}
