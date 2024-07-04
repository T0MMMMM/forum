package forum

import (
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/websocket/v2"
)

/*
This function allows you to send received messages directly to the pages.
*/
func (E *Engine) Websocket(c *websocket.Conn) {
	E.ConnectedUsers[c] = struct{}{}
	for {
		_, msg, err := c.ReadMessage()
		message := strings.Split(string(msg), ":") 
		if len(msg) > 0 {E.InsertMessages(message)}
		if err != nil {break}
		for usr := range E.ConnectedUsers {
			if usr != c {
				if (message[0] == "[TYPEPrivate]"){
					if err := usr.WriteMessage(websocket.TextMessage, []byte("[TYPEPrivate]:" + message[1]+":"+message[2]+":"+message[4]+":"+message[3])); err != nil {
						return
					}
				} else if (message[0] == "[TYPETopic]") {
					data := E.QuerySQL("SELECT id FROM answers")
					var id int
					var maxId int
					for data.Next() {
						data.Scan(&id)
						if (id > maxId) {
							maxId = id
						}
					}
					if err := usr.WriteMessage(websocket.TextMessage, []byte("[TYPETopic]:" + message[2] +":"+ message[3] +":"+ message[4] +":"+ message[5] +":"+  strconv.Itoa(maxId))); err != nil {
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

/*
This function allows you to update the database based on messages sent.
*/
func (E *Engine) InsertMessages(message []string) {
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
