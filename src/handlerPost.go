package forum

import (
	"strconv"
	"time"
	"github.com/gofiber/fiber/v2"
)

/*
This function allows the user to connect by entering their username and password.
If they do not match a User in the database, it does not log in and stays on the page, otherwise it logs in and is redigested on the main page
*/
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

/*
This function allows the user to register on the site. Its information will be stored in the database. 
If registration does not fail, the user is redirected to the main page and logged into their account
*/
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
			c.Redirect("/connexion")
			return nil
		} else if emailExist == email {
			E.CurrentData.ErrorMsg = "Cette adresse mail est déja utilisée, donc connecte toi stp "
			c.Redirect("/connexion")
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

/*
This function allows a user connected to their account to be able to create a new topic by choosing a category, 
a title, and a description. This information will be saved in the database.
*/
func (E *Engine) SubmitNewTopic(c *fiber.Ctx) error {
	E.GetCookieUser(c)
	categorieID := c.FormValue("categorie")
	title := c.FormValue("title")
	content := c.FormValue("content")
	if categorieID != "" && title != "" && content != "" && E.CurrentData.User.Username != "" {
		err := E.ExecuteSQL("INSERT INTO topics (categoryID, userID, title, content, created_at, status, visible, like, dislike) VALUES ('" + categorieID + "', '" + strconv.Itoa(E.CurrentData.User.Id) + "', '" + E.filterMsg(title) + "', '" + E.filterMsg(content) + "', '" + time.Now().String()[:19] + "', '" + "unsolved" + "', '" + "true" + "', '" + "0" + "', '" + "0" + "')")
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
	E.GetCookieUser(c)
	E.SetCookieTopic(E.StrToInt(TopicID), c)
	E.CurrentData.Topic = E.FindTopicByID(E.StrToInt(TopicID))
	c.Redirect("/")
	return nil
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
	E.SetCookieTopic(0, c)
	E.SetCookieCategory(categoryButton, c)
	c.Redirect("/")
	return c.Render("index", E.CurrentData)
}

func (E *Engine) SubmitSearch(c *fiber.Ctx) error {
	searchButton := c.FormValue("search")
	if searchButton != "" {E.CurrentData.CurrentSearch = searchButton}
	E.SetCookieTopic(0, c)
	E.SetCookieSearch(E.CurrentData.CurrentSearch, c)
	c.Redirect("/")
	return c.Render("index", E.CurrentData)
}

func (E *Engine) SubmitResetSearch(c *fiber.Ctx) error {
	E.CurrentData.CurrentSearch = ""
	E.SetCookieSearch(E.CurrentData.CurrentSearch, c)
	c.Redirect("/")
	return c.Render("index", E.CurrentData)
}

func (E *Engine) SubmitChangeUsername(c *fiber.Ctx) error {
	E.GetCookieUser(c)
	usernameButton := c.FormValue("username")
	var username string
	if usernameButton != "" {
		data := E.QuerySQL("SELECT username FROM users")
		for data.Next() {
			data.Scan(&username)
			if username == usernameButton {
				E.CurrentData.ErrorMsg = "Username already used"
				c.Redirect("/edit_profil")
				return nil
			}
		}
		E.ExecuteSQL("UPDATE users SET username = '" + usernameButton + "' WHERE id = " + strconv.Itoa(E.CurrentData.User.Id) + ";")
		c.Redirect("/view_profil")
		return nil
	}
	E.CurrentData.ErrorMsg = "Please provide a valid name"
	c.Redirect("/")
	return nil
}


func (E *Engine) SubmitChangePictureProfile(c *fiber.Ctx) error {
	picture := c.FormValue("picture")
	E.GetCookieUser(c)
	E.CurrentData.User.ProfilePicture = picture
	E.ExecuteSQL("UPDATE users SET profile_picture = '" + picture + "' WHERE id = " + strconv.Itoa(E.CurrentData.User.Id) + ";")
	c.Redirect("/")
	return nil
}

func (E *Engine) SubmitBackTopics(c *fiber.Ctx) error {
	E.GetCookieUser(c)
	E.SetCookieTopic(0, c)
	c.Redirect("/")
	return nil
}

func (E *Engine) SubmitRemoveTopic(c *fiber.Ctx) error {
	E.GetCookieUser(c)
	TopicID := c.FormValue("TopicID")
	E.ExecuteSQL("DELETE FROM topics WHERE id = " + TopicID + ";")
	E.SetCookieTopic(0, c)
	c.Redirect("/")
	return nil
}

func (E *Engine) ValidateAnswer(c *fiber.Ctx) error {
	E.GetCookieUser(c)
	button := c.FormValue("button")
	topicId := E.FindAnswerByID(E.StrToInt(button)).TopicID
	topic := E.FindTopicByID(topicId)
	var id int
	var status string
	var validate bool = false 
	data := E.QuerySQL("SELECT id, status FROM answers")
	for data.Next() {
		data.Scan(&id, &status)
		if E.StrToInt(button) == id && status == "solved" {
			E.ExecuteSQL("UPDATE topics SET status = 'unsolved' WHERE id = " + strconv.Itoa(topic.Id) + ";")
			E.ExecuteSQL("UPDATE answers SET status = 'unsolved' WHERE id = " + button + ";")
		} else if E.StrToInt(button) == id {
			E.ExecuteSQL("UPDATE topics SET status = 'solved' WHERE id = " + strconv.Itoa(topic.Id) + ";")
			E.ExecuteSQL("UPDATE answers SET status = 'solved' WHERE id = " + button + ";")
			validate = true
		}
	}
	data2 := E.QuerySQL("SELECT id, status FROM answers")
	for data2.Next() {
		data2.Scan(&id, &status)
		if E.StrToInt(button) != id && status == "solved" && validate {
			E.ExecuteSQL("UPDATE answers SET status = 'unsolved' WHERE id = " + strconv.Itoa(id) + ";")
		}
	}
	c.Redirect("/")
	return nil
}


