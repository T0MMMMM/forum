package forum

import (
	"strings"
	"time"

	"github.com/gofiber/websocket/v2"
)


func (E *Engine) Websocket(c *websocket.Conn) {
	E.ConnectedUsers[c] = struct{}{}
	for {
		_, msg, err := c.ReadMessage()
		message := strings.Split(string(msg), ":") 
		if len(msg) > 0 {
			if (message[0] == "[TYPEPrivate]"){
				E.ExecuteSQL("INSERT INTO messages (senderID, recipientID, content, created_at) VALUES ( '" + message[1] + "' , '" + message[3] + "', '" + E.filterMsg(message[4]) + "', '" + time.Now().String()[:19] + "')")
			} else if (message[0] == "[TYPETopic]") {
				E.ExecuteSQL("INSERT INTO answers (topicID, userID, content, created_at) VALUES ( '" + message[4] + "' , '" + message[1] + "', '" + E.filterMsg(message[3]) + "', '" + time.Now().String()[:19] + "')")
			} else if (message[0] == "[TYPELike]") {
				E.ExecuteSQL("INSERT INTO topicsLikes (topicID, userID) VALUES ( '" + message[2] + "' , '" + message[1] + "')")
				E.ExecuteSQL("UPDATE topics SET like = like + 1 WHERE id = " + message[2])
			} else if (message[0] == "[TYPEDislike]") {
				E.ExecuteSQL("INSERT INTO topicsDislikes (topicID, userID) VALUES ( '" + message[2] + "' , '" + message[1] + "')")
				E.ExecuteSQL("UPDATE topics SET dislike = dislike + 1 WHERE id = " + message[2])
			} else if (message[0] == "[TYPERemoveLike]") {
				E.ExecuteSQL("DELETE FROM topicsLikes WHERE topicID = " + message[2] + " AND userID = " + message[1] + ";")
				E.ExecuteSQL("UPDATE topics SET like = like - 1 WHERE id = " + message[2])
			} else if (message[0] == "[TYPERemoveDislike]") {
				E.ExecuteSQL("DELETE FROM topicsDislikes WHERE topicID = " + message[2] + " AND userID = " + message[1] + ";")
				E.ExecuteSQL("UPDATE topics SET dislike = dislike - 1 WHERE id = " + message[2])
			}
		}
		if err != nil {
			break
		}
		for usr := range E.ConnectedUsers {
			if usr != c {
				if (message[0] == "[TYPEPrivate]"){ // 0 = [TYPEPrivate] / 1 = id / 2 = username / 3 = recipientId / 4 = msg 
					if err := usr.WriteMessage(websocket.TextMessage, []byte("[TYPEPrivate]:" + message[1]+":"+message[2]+":"+message[4]+":"+message[3])); err != nil {
						return
					}
				} else if (message[0] == "[TYPETopic]") { // 0 = "[TYPETopic]"  / 1 = id / 2 = user / 3 = msg  / 4 = topicId
					if err := usr.WriteMessage(websocket.TextMessage, []byte("[TYPETopic]:" + message[2]+":"+message[3]+":"+message[4])); err != nil {
						return
					}
				} else {
					if err := usr.WriteMessage(websocket.TextMessage, []byte(message[0])); err != nil {
						return
					}
				}
			}
		}
	}
	delete(E.ConnectedUsers, c)
	c.Close()
}
