package forum

import (
	"strconv"
	"strings"
)

func (E *Engine) StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func (E *Engine) filterMsg(msg string) string {
	return strings.ReplaceAll(msg, "'", "[[apostroph]]")
}

func (E *Engine) reversefilterMsg(msg string) string {
	return strings.ReplaceAll(msg, "[[apostroph]]", "'")
}

func (E *Engine) UsersMessages() {
	E.CurrentData.Users = []User{}
	data := E.QuerySQL("SELECT id FROM users")
	var id int
	for data.Next() {
		data.Scan(&id)
		user := E.FindUserByID(id);
		user.Messages = E.ScanMessages(user)
		E.CurrentData.Users = append(E.CurrentData.Users, user)
	}
}

func (E *Engine) ScanMessages(user User) []Message {
	data := E.QuerySQL("SELECT id FROM messages")
	var id int
	var list []Message
	for data.Next() {
		data.Scan(&id)
		msg := E.FindMessageByID(id)
		if ((user.Id == msg.Recipient.Id && msg.Sender.Id == E.CurrentData.User.Id) || (user.Id == msg.Sender.Id && msg.Recipient.Id == E.CurrentData.User.Id)) {
			list = append(list, msg)
		}
	}
	return list
}

func (E *Engine) SetTopics() {
	E.CurrentData.Topics = []Topic{}
	data := E.QuerySQL("SELECT id FROM topics")
	var id int
	for data.Next() {
		data.Scan(&id)
		topic := E.FindTopicByID(id)
		topic.Liked = E.SetLikedAndDisliked("topicsLikes", topic)
		topic.Disliked = E.SetLikedAndDisliked("topicsDislike", topic)
		E.CurrentData.Topics = append(E.CurrentData.Topics, topic)
	}
}

func (E *Engine) SetLikedAndDisliked(table string, topic Topic) bool {
	data := E.QuerySQL("SELECT userID FROM topicsLikes WHERE topicID = " + strconv.Itoa(topic.Id))
	var userID int
	for data.Next() {
		data.Scan(&userID)
		if (E.CurrentData.User.Id == userID) {
			return true
		}
	}
	return false
}
