package forum

import (
	"fmt"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func (E *Engine) SubmitConnexion(c *fiber.Ctx) error {
	username := c.FormValue("username")
	pwd := c.FormValue("pwd")
	if username != "" && pwd != "" {
		data := E.QuerySQL("SELECT id, username, password, email, created_at FROM users")
		var usernameRnd string
		var passwordRnd string
		var email string
		var id int
		var created_at string
		for data.Next() {
			data.Scan(&id, &usernameRnd, &passwordRnd, &email, &created_at)
			if usernameRnd == username && passwordRnd == pwd {
				E.SetCookieUser(id, c)
				E.CurrentData.ErrorMsg = ""
				c.Redirect("/")
				return c.SendString("0")
			}
		}
	}
	E.CurrentData.ErrorMsg = "Nom d'utilisateur ou mail incorect"
	c.Redirect("/connexion")
	return c.SendString("2")
}

func (E *Engine) SubmitRegister(c *fiber.Ctx) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	pwd := c.FormValue("pwd")
	fmt.Println("d")
	fmt.Println(username)
	
	var usernameExist string
	var emailExist string

	data := E.QuerySQL("SELECT username, email FROM users")
	for data.Next() {
		data.Scan(&usernameExist, &emailExist)
		if usernameExist == username {
			E.CurrentData.ErrorMsg = "Vous utilisez un nom d'utilisateur déja existant, donc rajoute full _"
			c.Redirect("/register")
			return nil
		} else if emailExist == email {
			E.CurrentData.ErrorMsg = "Cette adresse mail est déja utilisée, donc connecte toi stp "
			c.Redirect("/register")
			return nil
		}
	}
	if username != "" && pwd != "" && email != "" {
		err := E.ExecuteSQL("INSERT INTO users (username, password, email) VALUES ('" + username + "', '" + pwd + "', '" + email + "')")
		if err != nil {
			E.CurrentData.ErrorMsg = "Erreur de base de données, donc rien à voir avec vous, réessaie plus tard"
			c.Redirect("/connexion")
			return c.SendString("1")
		} else {
			var count int
			data :=  E.QuerySQL("SELECT count(*) FROM users")
			for data.Next() {
				err:= data.Scan(&count)
				if (err != nil) {panic(err)}
			}
			E.SetCookieUser(count, c)
			c.Redirect("/")
		}
	} else {
		c.Redirect("/connexion")
	}
	c.Redirect("/")
	return c.SendString("0")
}

func (E *Engine) SubmitNewTopic(c *fiber.Ctx) error {
	categorieID := c.FormValue("categorie")
	title := c.FormValue("title")
	content := c.FormValue("content")

	if categorieID != "" && title != "" && content != "" && E.CurrentData.User.Username != "" {
		err := E.ExecuteSQL("INSERT INTO topics (categoryID, userID, title, content, created_at, status, visible, like, dislike) VALUES ('" + categorieID + "', '" + strconv.Itoa(E.CurrentData.User.Id) + "', '" + title + "', '" + content + "', '" + "2006-01-02 15:04:05" + "', '" + "unsolved" + "', '" + "true" + "', '" + "0" + "', '" + "0" + "')")
		if err != nil {
			E.CurrentData.ErrorMsg = "Erreur de base de données, donc rien à voir avec vous, réessaie plus tard"
			c.Redirect("/new-topic")
			return c.SendString("2")
		}
	} else {
		E.CurrentData.ErrorMsg = "Champs obligatoires / vous n'etes pas connectés gros nul"
		c.Redirect("/new-topic")
		return c.SendString("2")
	}

	c.Redirect("/")
	return c.SendString("0")
}