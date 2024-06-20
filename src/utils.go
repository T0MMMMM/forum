package forum

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	//"strconv"
	"strings"
)

func (E *Engine) DataBaseCreation() {
	sqlFile, err := ioutil.ReadFile("./serv/sql/forum.sql")
	if err != nil {
		fmt.Printf("Error reading SQL file: %v", err)
	}

	if err != nil {
		fmt.Printf("Error connecting to the database: %v", err)
	}

	commands := splitSQLCommands(string(sqlFile))

	for _, cmd := range commands {
		_, err := E.DataBase.Exec(cmd)
		if err != nil {
			log.Fatalf("Erreur d'exécution de la commande SQL: %v", err)
		}
	}
	_, err = E.DataBase.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		log.Fatal(err)
	}

	// Catégories Creation

	//E.ExecuteSQL("INSERT INTO `categories` (name, description) VALUES ('Developpement', 'Developpement informatique')")
	//INSERT INTO categories (name, description) VALUES ('Developpement', 'Developpement informatique');
	//INSERT INTO categories (name, description) VALUES ('Cinéma', 'Cinéma films');
	//DELETE FROM categories WHERE id > 2;

	E.DataBase.Close()
}

func splitSQLCommands(file string) []string {
	list := strings.Split(file, ";")
	list = list[:len(list)-1]
	for i := 0; i < len(list); i++ {
		list[i] += ";"
	}
	return list
}

func (E *Engine) ExecuteSQL(cmd string) error {
	E.DataBase, _ = sql.Open("sqlite", "./serv/data/data.db")
	_, err := E.DataBase.Exec(cmd)
	if err != nil {
		log.Fatalf("Erreur d'exécution de la commande SQL: %v", err)
		return err
	}
	E.DataBase.Close()
	return nil
}

func (E *Engine) QuerySQL(cmd string) *sql.Rows {
	E.DataBase, _ = sql.Open("sqlite", "./serv/data/data.db")
	data, err := E.DataBase.Query(cmd)
	if err != nil {
		log.Fatalf("Erreur d'exécution de la commande SQL: %v", err)
	}
	E.DataBase.Close()
	return data
}

func (E *Engine) FindTopicByID(TopicID int) Topic {

	data := E.QuerySQL("SELECT id, categoryID, userID, title, content, created_at, status, visible, like, dislike FROM topics WHERE id = " + strconv.Itoa(TopicID))

	var (
		id         int
		title      string
		categoryID int
		userID     int
		content    string
		created_at string
		status     string
		visible    bool
		like       int
		dislike    int
	)

	for data.Next() {
		data.Scan(&id, &categoryID, &userID, &title, &content, &created_at, &status, &visible, &like, &dislike)
	}

	return Topic{
		Id:        id,
		Title:     title,
		Content:   content,
		Category:  E.FindCategoryByID(categoryID),
		User:      E.FindUserByID(userID),
		CreatedAt: created_at,
		Status:    status,
		Visible:   visible,
		Like:      like,
		Dislike:   dislike,
		Answers:   E.FindAnswersByTopicID(TopicID),
	}
}

func (E *Engine) FindCategoryByID(CategoryID int) Category {
	data := E.QuerySQL("SELECT id, name, description FROM Categories WHERE id = " + strconv.Itoa(CategoryID))

	var (
		id          int
		name        string
		description string
	)

	for data.Next() {
		data.Scan(&id, &name, &description)
	}

	return Category{
		Id:          id,
		Name:        name,
		Description: description,
	}
}

func (E *Engine) FindUserByID(UserID int) User {
	data := E.QuerySQL("SELECT id, email, username, password, created_at FROM users WHERE id = " + strconv.Itoa(UserID))
	var (
		id         int
		email      string
		username   string
		password   string
		created_at string
	)
	for data.Next() {
		data.Scan(&id, &email, &username, &password, &created_at)
	}
	return User{
		Id:        id,
		Username:  username,
		Email:     email,
		CreatedAt: created_at,
	}
}

func (E *Engine) FindAnswersByTopicID(TopicID int) []Answer {

	data := E.QuerySQL("SELECT id, userID, content, created_at, status, visible, like, dislike FROM answers WHERE topicID = " + strconv.Itoa(TopicID))

	var (
		answers    []Answer
		id         int
		userID     int
		content    string
		created_at string
		status     string
		visible    bool
		like       int
		dislike    int
	)

	for data.Next() {
		data.Scan(&id, &userID, &content, &created_at, &status, &visible, &like, &dislike)
		answers = append(answers, Answer{Id: id, User: E.FindUserByID(userID), Content: content, CreatedAt: created_at, Status: status, Visible: visible, Like: like, Dislike: dislike})
	}
	return answers
}

func (E *Engine) GetCookieUser(c *fiber.Ctx) {
	if (c.Cookies("UserId", "false") == "false") {
		E.CurrentData.User = E.FindUserByID(E.StrToInt(c.Cookies("UserID")))
	}
}

func (E *Engine) GetCookieTopic(c *fiber.Ctx) {
	E.CurrentData.Topic = E.FindTopicByID(E.StrToInt(c.Cookies("TopicID")))
}

func (E *Engine) SetCookieUser(userID int, c *fiber.Ctx) {
	cookieUser := new(fiber.Cookie)
	cookieUser.Name = "UserID"
	cookieUser.Value = strconv.Itoa(userID)
	cookieUser.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookieUser)
}

func (E *Engine) SetCookieTopic(topicID int, c *fiber.Ctx) {
	cookieTopic := new(fiber.Cookie)
	cookieTopic.Name = "TopicID"
	cookieTopic.Value = strconv.Itoa(topicID)
	cookieTopic.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookieTopic)
}

func (E *Engine) StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

