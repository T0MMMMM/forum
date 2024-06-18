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
	defer func() {E.CurrentData.CurrentErrorMsg = ""} ()
	return c.Render("connexion", E.CurrentData)
}

func (E *Engine) Register(c *fiber.Ctx) error {
	defer func() {E.CurrentData.CurrentErrorMsg = ""} ()
	return c.Render("register", E.CurrentData)
}

func (E *Engine) NewTopic(c *fiber.Ctx) error {
	defer func() {E.CurrentData.CurrentErrorMsg = ""} ()
	return c.Render("new-topic", E.CurrentData)
}

func (E *Engine) Topic(c *fiber.Ctx) error {

	TopicID := c.FormValue("TopicID")

	defer func() {E.CurrentData.CurrentErrorMsg = ""} ()
	return c.Render("topic", E.FindTopicByID(E.StrToInt(TopicID)))
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
				E.CurrentData.CurrentErrorMsg = ""
				c.Redirect("/")
				return c.SendString("0")
			}
		}
	}
	E.CurrentData.CurrentErrorMsg = "Nom d'utilisateur ou mail incorect"
	c.Redirect("/connexion")
	return c.SendString("1")
}

func (E *Engine) SubmitRegister(c *fiber.Ctx) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	pwd := c.FormValue("pwd")

	var usernameExist string
	var emailExist string

	data := E.QuerySQL("SELECT username, email FROM users")
	for data.Next() {
		data.Scan(&usernameExist, &emailExist)
		if usernameExist == username  {
			E.CurrentData.CurrentErrorMsg = "Vous utilisez un nom d'utilisateur déja existant, donc rajoute full _"
			c.Redirect("/register")
			return nil
		} else if emailExist == email {
			E.CurrentData.CurrentErrorMsg = "Cette adresse mail est déja utilisée, donc connecte toi stp "
			c.Redirect("/register")
			return nil
		}
	}
	if (username != "" && pwd != "" && email != "") {
		err := E.ExecuteSQL("INSERT INTO users (username, password, email) VALUES ('" + username + "', '" + pwd + "', '" + email + "')")
		if (err != nil) {
			E.CurrentData.CurrentErrorMsg = "Erreur de base de données, donc rien à voir avec vous, réessaie plus tard"
			c.Redirect("/register")
			return c.SendString("1")
		}
	}
	c.Redirect("/")
	return c.SendString("0")
}



func (E *Engine) SubmitNewTopic(c *fiber.Ctx) error {
	categorieID := c.FormValue("categorie")
	title := c.FormValue("title")
	content := c.FormValue("content")

	if (categorieID != "" && title != "" && content != "" && E.CurrentUser.Username != "") {
		err := E.ExecuteSQL("INSERT INTO topics (categoryID, userID, title, content, created_at, status, visible, like, dislike) VALUES ('" + categorieID + "', '" + strconv.Itoa(E.CurrentUser.Id) + "', '" + title + "', '" + content + "', '" + "2006-01-02 15:04:05" + "', '" + "unsolved" + "', '" + "true" + "', '" + "0" + "', '" + "0" + "')")
		if (err != nil) {
			E.CurrentData.CurrentErrorMsg = "Erreur de base de données, donc rien à voir avec vous, réessaie plus tard"
			c.Redirect("/new-topic")
			return c.SendString("1")
		}
	} else {
		E.CurrentData.CurrentErrorMsg = "Champs obligatoires / vous n'etes pas connectés gros nul"
		c.Redirect("/new-topic")
		return c.SendString("1")
	}

	c.Redirect("/")
	return c.SendString("0")
}



func (E *Engine) Websocket(c *websocket.Conn) {

    E.ConnectedUsers[c] = struct{}{}

        for {
            _, msg, err := c.ReadMessage()
			if (len(msg) > 0) {
				E.ExecuteSQL("INSERT INTO answers (userID, content) VALUES ('" + strconv.Itoa(E.CurrentUser.Id) + "', '" + string(msg[len(E.CurrentUser.Username)+3:]) + "')")
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
