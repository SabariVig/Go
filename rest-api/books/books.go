package books

import "github.com/gofiber/fiber"


type book struct {
	name string
}

func GetBooks(c *fiber.Ctx){
	c.Send(book{"Hawk"})
}
func GetBook(c *fiber.Ctx){
	c.Send("Single Book")
}


func NewBook(c *fiber.Ctx){
	c.Send("Book Updated")
}
func DeleteBook(c *fiber.Ctx){
	c.Send("Book Deleted")
}

