package forum

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
	_ "modernc.org/sqlite"
)

type Engine struct {
	Port           string
	ConnectedUsers map[*websocket.Conn]struct{}
	DataBase       *sql.DB
	CurrentUser    User
	CurrentData    Data
}

type Data struct {
	CurrentErrorMsg string
	Categories      []Category
}

type Category struct {
	Id          int
	Name        string
	Description string
}

type User struct {
	Id       int
	Username string
	Email    string
	Password string
	CreatedAt string
}

type Topic struct {
	Id int
	Title string
	Category Category
	User User
	Content string
	CreatedAt string
	Status string
	Visible bool
	Like int
	Dislike int

	Answers []Answer
}

type Answer struct {
	Id int
	User User
	Content string
	CreatedAt string
	Status string
	Visible bool
	Like int
	Dislike int
}

func (E *Engine) Init() {
	//rand.Seed(time.Now().UnixNano())

	E.CurrentUser = User{
		Username: "",
		Email:    "",
		Password: "",
	}

	E.CurrentData = Data{
		CurrentErrorMsg: "",
	}

	E.DataBase, _ = sql.Open("sqlite", "./serv/data/data.db")

	E.DataBaseCreation()
	E.Port = ":8080"

	data := E.QuerySQL("SELECT id, name, description FROM categories")
	var id int
	var name string
	var description string
	for data.Next() {
		data.Scan(&id, &name, &description)
		E.CurrentData.Categories = append(E.CurrentData.Categories, Category{Id: id, Name: name, Description: description})
	}

}

func (E *Engine) Run() {
	E.Init()
	engine := html.New("./serv/html", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/serv", "./serv")

	E.ConnectedUsers = make(map[*websocket.Conn]struct{})

	// WebSocket route for user communication
	app.Get("/ws", websocket.New(E.Websocket))

	app.Get("/", E.Index)

	app.Get("/connexion", E.Connexion)
	app.Get("/register", E.Register)
	app.Get("/new-topic", E.NewTopic)
	app.Get("/topic", E.Topic)


	app.Post("/submit_connexion", E.SubmitConnexion)
	app.Post("/submit_register", E.SubmitRegister)
	app.Post("/submit_new-topic", E.SubmitNewTopic)

	app.Listen(E.Port)
}
