package forum

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
	_ "modernc.org/sqlite"
)

func (E *Engine) Run() {
	E.Init()
	engine := html.New("./serv/html", ".html")
	app := fiber.New(fiber.Config{
		Views:     engine,
		Immutable: true,
	})
	E.ConnectedUsers = make(map[*websocket.Conn]struct{})
	app.Static("/serv", "./serv")
	app.Get("/ws", websocket.New(E.Websocket))

	app.Get("/", E.Index)
	app.Get("/connexion", E.Connexion)
	app.Get("/new-topic", E.NewTopic)
	app.Get("/topic", E.Topic)

	app.Post("/submit_topic", E.SubmitTopic)
	app.Post("/submit_connexion", E.SubmitConnexion)
	app.Post("/submit_register", E.SubmitRegister)
	app.Post("/submit_new-topic", E.SubmitNewTopic)

	app.Listen(E.Port)
}