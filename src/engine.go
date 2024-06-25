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
	Users      []User
	UserSearch UserSearch
	CurrentCategory string
	CurrentSearch string
}

type Category struct {
	Id          int
	Name        string
	Description string
}

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	CreatedAt string
	ProfilePicture string
	Messages  []Message
}

type UserSearch struct {
	Username  string
	Email     string
	CreatedAt string
	ProfilePicture string
	Topics    []Topic
	AnswersTopic []AnswerTopic
}

type AnswerTopic struct {
	Answer   Answer
	Topic    Topic
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
	Liked	  bool
	Disliked  bool
	Answers []Answer
}

type Message struct {
	Id 		int
	Sender User
	Recipient User
	Content string
	Visible bool
	CreatedAt string
}

type Answer struct {
	Id        int
	User      User
	TopicID   int
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
	E.CurrentData.CurrentCategory = ""
	E.CurrentData.CurrentSearch = ""
	E.CurrentData.UserSearch = UserSearch{}
	E.InitDescriptions()
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

