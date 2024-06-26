package forum

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (E *Engine) SubmitConnexion(c *fiber.Ctx) error {
	username := c.FormValue("username")
	pwd := c.FormValue("pwd")
	if username != "" && pwd != "" {
		data := E.QuerySQL("SELECT id, username, password, email FROM users")
		var usernameRnd string
		var passwordRnd string
		var email string
		var id int
		for data.Next() {
			data.Scan(&id, &usernameRnd, &passwordRnd, &email)
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
		err := E.ExecuteSQL("INSERT INTO users (username, password, email, created_at, profile_picture) VALUES ('" + username + "', '" + pwd + "', '" + email + "', '" + time.Now().String()[:19] + "', '" + "serv/assets/pictures/default.jpg" + "')")
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
		err := E.ExecuteSQL("INSERT INTO topics (categoryID, userID, title, content, created_at, status, visible, like, dislike) VALUES ('" + categorieID + "', '" + strconv.Itoa(E.CurrentData.User.Id) + "', '" + title + "', '" + content + "', '" + time.Now().String()[:19] + "', '" + "unsolved" + "', '" + "true" + "', '" + "0" + "', '" + "0" + "')")
		E.CurrentData.Topics = append(E.CurrentData.Topics, E.FindTopicByID(len(E.CurrentData.Topics)+1))
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

func (E *Engine) SubmitTopic(c *fiber.Ctx) error {
	TopicID := c.FormValue("TopicID")
	//E.SetCookieTopic(E.StrToInt(TopicID), c)
	E.GetCookieUser(c)
	E.CurrentData.Topic = E.FindTopicByID(E.StrToInt(TopicID))
	//E.GetCookieTopic(c)
	defer func() { E.CurrentData.ErrorMsg = "" }()
	return c.Render("topic", E.CurrentData)
}

func (E *Engine) SubmitUser(c *fiber.Ctx) error {
	UserID := c.FormValue("user")
	E.GetCookieUser(c)
	E.CurrentData.UserSearch = E.CreateUserSearch(E.StrToInt(UserID))
	defer func() { E.CurrentData.ErrorMsg = "" }()
	return c.Render("userSearch", E.CurrentData)
}


func (E *Engine) SubmitChoseCategory(c *fiber.Ctx) error {
	categoryButton := c.FormValue("category")
	E.CurrentData.CurrentCategory = categoryButton
	c.Redirect("/")
	return c.Render("index", E.CurrentData)
}

func (E *Engine) SubmitSearch(c *fiber.Ctx) error {
	searchButton := c.FormValue("search")
	if searchButton != "" {E.CurrentData.CurrentSearch = searchButton}
	c.Redirect("/")
	return c.Render("index", E.CurrentData)
}

func (E *Engine) SubmitResetSearch(c *fiber.Ctx) error {
	E.CurrentData.CurrentSearch = ""
	c.Redirect("/")
	return c.Render("index", E.CurrentData)
}

func (E *Engine) SubmitChangeUsername(c *fiber.Ctx) error {
	usernameButton := c.FormValue("username")
	var username string
	if usernameButton != "" {
		data := E.QuerySQL("SELECT username FROM users")
		for data.Next() {
			data.Scan(&username)
			if username == usernameButton {
				E.CurrentData.ErrorMsg = "Username already used"
				E.GetCookieUser(c)
				c.Redirect("/edit_profil")
				return nil
			}
		}
		E.ExecuteSQL("UPDATE users SET username = '" + usernameButton + "' WHERE id = " + strconv.Itoa(E.CurrentData.User.Id) + ";")
		E.GetCookieUser(c)
		c.Redirect("/edit_profil")
		return nil
	}
	E.CurrentData.ErrorMsg = "Please provide a valid name"
	E.GetCookieUser(c)
	c.Redirect("/edit_profil")
	return c.SendString("1")
}


func (E *Engine) SubmitChangePictureProfile(c *fiber.Ctx) error {
	picture := c.FormValue("picture")
	E.CurrentData.User.ProfilePicture = picture
	E.ExecuteSQL("UPDATE users SET profile_picture = '" + picture + "' WHERE id = " + strconv.Itoa(E.CurrentData.User.Id) + ";")
	E.GetCookieUser(c)
	c.Redirect("/edit_profil")
	return c.SendString("1")
}




