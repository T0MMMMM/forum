package forum

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
)

type Engine struct {
	Port string
    ConnectedUsers map[*websocket.Conn]struct{}
    DataBase *sql.DB
}

func (E *Engine) Init() {
	//rand.Seed(time.Now().UnixNano())

    E.DataBase, _ = sql.Open("sqlite3", "./serv/data/data.db")

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

    app.Post("/submit", E.Submit)

	app.Listen(E.Port)
}
