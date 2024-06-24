package forum

import (
	"github.com/gofiber/fiber/v2"
)


func (E *Engine) Index(c *fiber.Ctx) error {
	E.GetCookieUser(c)
	E.UsersMessages()
	E.SetTopics()
	defer func() { E.CurrentData.ErrorMsg = "" }()
	return c.Render("index", E.CurrentData)
}

func (E *Engine) Connexion(c *fiber.Ctx) error {
	//E.GetCookieUser(c)
	defer func() { E.CurrentData.ErrorMsg = "" }()
	return c.Render("connexion", E.CurrentData)
}

func (E *Engine) NewTopic(c *fiber.Ctx) error {
	E.GetCookieUser(c)
	E.UsersMessages()
	defer func() { E.CurrentData.ErrorMsg = "" }()
	return c.Render("new-topic", E.CurrentData)
}

func (E *Engine) Topic(c *fiber.Ctx) error {
	E.GetCookieUser(c)
	E.UsersMessages()
	defer func() { E.CurrentData.ErrorMsg = "" }()
	return c.Render("topic", E.CurrentData)
}


