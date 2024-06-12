package forum

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	//"github.com/gofiber/websocket/v2"
)

type Engine struct {
	Port string
}

func (E *Engine) Init() {
	//rand.Seed(time.Now().UnixNano())
	E.Port = ":8080"
}

func (E *Engine) Run() {
	E.Init()
	engine := html.New("./serv/html", ".html")
	
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/serv", "./serv")
	app.Get("/", E.Index)
	app.Listen(E.Port)
}
