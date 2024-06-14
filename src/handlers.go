package forum

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func (E *Engine) Index(c *fiber.Ctx) error {
	return c.Render("index", E.CurrentUser)
}

func (E *Engine) Connexion(c *fiber.Ctx) error {
	return c.Render("connexion", nil)
}

func (E *Engine) Register(c *fiber.Ctx) error {
	return c.Render("register", nil)
}

func (E *Engine) SubmitConnexion(c *fiber.Ctx) error {
	username := c.FormValue("username")
	pwd := c.FormValue("pwd")
	
	if (username != "" && pwd != "") {
		data := E.QuerySQL("SELECT id, username, password FROM users")
		var usernameRnd string
		var passwordRnd string
		var id int
		for data.Next() {
			data.Scan(&id, &usernameRnd, &passwordRnd)
			if usernameRnd == username && passwordRnd == pwd {
				E.CurrentUser = User{
					Id : id,
					Username: username,
					Email:  "",
					Password: pwd,
				}
				c.Redirect("/")
				return c.SendString("0")
			}
		}
	}
	c.Redirect("/connexion")
	return c.SendString("1")
}

func (E *Engine) SubmitRegister(c *fiber.Ctx) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	pwd := c.FormValue("pwd")
	if (username != "" && pwd != "" && email != "") {
		err := E.ExecuteSQL("INSERT INTO users (username, password, email) VALUES ('" + username + "', '" + pwd + "', '" + email + "')")
		if (err != nil) {
			c.Redirect("/register")
			return c.SendString("1")
		}
	}
	c.Redirect("/")
	return c.SendString("0")
}

func (E *Engine) Websocket(c *websocket.Conn) {

    E.ConnectedUsers[c] = struct{}{}

        for {
            _, msg, err := c.ReadMessage()
			if (len(msg) > 0) {
				E.ExecuteSQL("INSERT INTO Posts (userID, content) VALUES ('" + strconv.Itoa(E.CurrentUser.Id) + "', '" + string(msg[len(E.CurrentUser.Username)+3:]) + "')")
			}
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
