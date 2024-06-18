package forum

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
    _ "modernc.org/sqlite"
    "database/sql"
)

type Engine struct {
	Port string
    ConnectedUsers map[*websocket.Conn]struct{}
    DataBase *sql.DB
	CurrentUser User
	CurrentData Data 
}

type Data struct {
	CurrentErrorMsg string
}


type User struct {
	Id int
	Username string
	Email string
	Password string
}

func (E *Engine) Init() {
	//rand.Seed(time.Now().UnixNano())

	E.CurrentUser = User{
		Username: "",
		Email:  "",
		Password: "",
	}
	E.CurrentData = Data{
		CurrentErrorMsg : "",
	}

    E.DataBase, _ = sql.Open("sqlite", "./serv/data/data.db")

	E.DataBaseCreation()
	E.Port = ":8080"

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


    app.Post("/submit_connexion", E.SubmitConnexion)
	app.Post("/submit_register", E.SubmitRegister)


	app.Listen(E.Port)
}
