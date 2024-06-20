package forum

import (
	"database/sql"
	"github.com/gofiber/websocket/v2"
	_ "modernc.org/sqlite"
)

type Engine struct {
	Port           string
	ConnectedUsers map[*websocket.Conn]struct{}
	DataBase       *sql.DB
	CurrentData    Data
}

type Data struct {
	Topic      Topic
	User       User
	ErrorMsg   string
	Categories []Category
	Topics     []Topic
}

type Category struct {
	Id          int
	Name        string
	Description string
}

type User struct {
	Id        int    `cookie:"id"`
	Username  string `cookie:"username"`
	Email     string `cookie:"email"`
	Password  string `cookie:"password"`
	CreatedAt string `cookie:"created_at"`
}

type Topic struct {
	Id        int
	Title     string
	Category  Category
	User      User
	Content   string
	CreatedAt string
	Status    string
	Visible   bool
	Like      int
	Dislike   int
	Answers []Answer
}

type Answer struct {
	Id        int
	User      User
	Content   string
	CreatedAt string
	Status    string
	Visible   bool
	Like      int
	Dislike   int
}



func (E *Engine) Init() {
	//rand.Seed(time.Now().UnixNano())
	E.CurrentData.User = User{}
	E.CurrentData = Data{ErrorMsg: "",}
	E.DataBase, _ = sql.Open("sqlite", "./serv/data/data.db")
	E.DataBaseCreation()
	E.Port = ":8080"
	E.InitDescriptions()
	E.InitTopics()
}



func (E *Engine) InitDescriptions() {
	data := E.QuerySQL("SELECT id, name, description FROM categories")
	var id int
	var name string
	var description string
	for data.Next() {
		data.Scan(&id, &name, &description)
		E.CurrentData.Categories = append(E.CurrentData.Categories, Category{Id: id, Name: name, Description: description})
	}
}

func (E *Engine) InitTopics() {
	data := E.QuerySQL("SELECT id, categoryID, userID, title, content, created_at, status, visible, like, dislike FROM topics")
	var id int
	var categoryID int
	var userID int
	var title string
	var content string
	var created_at string
	var status string
	var visible bool
	var like int
	var dislike int
	for data.Next() {
		data.Scan(&id, &categoryID, &userID, &title, &content, &created_at, &status, &visible, &like, &dislike)
		E.CurrentData.Topics = 
		append(E.CurrentData.Topics, Topic{Id: id, Category: E.FindCategoryByID(categoryID), 
			User: E.FindUserByID(userID), Title: title, Content: content, CreatedAt: created_at, Status: status,
		    Visible: visible, Like: like, Dislike: dislike,})
	}
}
