package forum

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
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

func (E *Engine) Connexion(c *fiber.Ctx) error {
	return c.Render("connexion", nil)
}


func (E *Engine) Submit(c *fiber.Ctx) error {
	username := c.FormValue("username")
	pwd := c.FormValue("pwd")

	if (username != "" && pwd != "") {
		fmt.Println(username)
		fmt.Println(" --- ")
		fmt.Println(pwd)
		
		//E.DataBase.Exec()
	}

	return c.SendString("c ok")
}

func (E *Engine) Websocket(c *websocket.Conn) {

    E.ConnectedUsers[c] = struct{}{}

        for {
            _, msg, err := c.ReadMessage()
            if err != nil {
                break
            }
            for usr := range E.ConnectedUsers {
                if usr != c {
                    if err := usr.WriteMessage(websocket.TextMessage, msg); err != nil {
                        return
                    }
                }
            }
        }
	delete(E.ConnectedUsers, c)
		
	c.Close()
}
