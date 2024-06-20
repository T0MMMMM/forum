package forum

import (
	"strings"
	"github.com/gofiber/websocket/v2"
)


func (E *Engine) Websocket(c *websocket.Conn) {
	E.ConnectedUsers[c] = struct{}{}
	for {
		_, msg, err := c.ReadMessage()
		message := strings.Split(string(msg), ":") // 0 = id / 1 = username / 2 = msg / 3 = topic ID
		if len(msg) > 0 {
			E.ExecuteSQL("INSERT INTO answers (topicID, userID, content) VALUES ( '" + message[3] + "' , '" + message[0] + "', '" + message[2] + "')")
		}
		if err != nil {
			break
		}
		for usr := range E.ConnectedUsers {
			if usr != c {
				if err := usr.WriteMessage(websocket.TextMessage, []byte(message[1]+":"+message[2]+":"+message[3])); err != nil {
					return
				}
			}
		}
	}
	delete(E.ConnectedUsers, c)
	c.Close()
}
