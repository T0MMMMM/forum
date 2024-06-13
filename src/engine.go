package forum

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
)

type Engine struct {
	Port string
}

func (E *Engine) Init() {
	//rand.Seed(time.Now().UnixNano())
	E.DataBase()
	E.Port = ":8080"
}

func (E *Engine) Run() {
	E.Init()
	engine := html.New("./serv/html", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/serv", "./serv")

	connectedUsers := make(map[*websocket.Conn]struct{})

    // WebSocket route for user communication
    app.Get("/ws", websocket.New(func(c *websocket.Conn) {
        // Add WebSocket connection to active connections map
        connectedUsers[c] = struct{}{}
		
        defer func() {
            // Remove WebSocket connection from active connections map
            delete(connectedUsers, c)
            // Close WebSocket connection
            c.Close()
        }()

        // Infinite loop to listen for incoming messages
        for {
            // Read message from client
            _, msg, err := c.ReadMessage()
            if err != nil {
                break
            }

            // Transmit message to all other connected users
            for usr := range connectedUsers {
                if usr != c {
                    if err := usr.WriteMessage(websocket.TextMessage, msg); err != nil {
                        return
                    }
                }
            }
        }
    }))

	app.Get("/", E.Index)
	app.Listen(E.Port)
}
