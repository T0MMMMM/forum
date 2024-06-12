package forum

import (
	"github.com/gofiber/fiber/v2"
)
	
type Person struct {
	Name string
	Age  int
}

func (E *Engine) Index(c *fiber.Ctx) error {
	person := Person{
		Name: "John Connor",
		Age:  30,
	}
	return c.Render("index", person)
}