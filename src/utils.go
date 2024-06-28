package forum

import (
	"strconv"
	"strings"
)

/*
This function transforms a string type variable into int. It is used to retransform ids
*/
func (E *Engine) StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

/*
this function replace apostrophes as these can cause problems in the database
*/
func (E *Engine) filterMsg(msg string) string {
	return strings.ReplaceAll(msg, "'", "[[apostroph]]")
}

/*
this function transforms converted apostrophes into real apostrophes
*/
func (E *Engine) reversefilterMsg(msg string) string {
	return strings.ReplaceAll(msg, "[[apostroph]]", "'")
}

/*
This function allows you to load messages between users. It is used on the main page
*/
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

/*
This function allows you to return the list of messages for a user if he participates in the message (sender or receiver)
*/
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

/*
This function allows you to generate topics on the main page. It takes into account user filters
*/
func (E *Engine) SetTopics() {
	E.CurrentData.Topics = []Topic{}
	data := E.QuerySQL("SELECT id FROM topics")
	var id int
	for data.Next() {
		data.Scan(&id)
		topic := E.FindTopicByID(id)
		topic.Liked = E.SetLikedAndDisliked("topicsLikes", topic)
		topic.Disliked = E.SetLikedAndDisliked("topicsDislikes", topic)
		if E.CurrentData.CurrentCategory == "" {
			if E.ContainsTxt(topic) {E.CurrentData.Topics = append(E.CurrentData.Topics, topic)}
		} else if E.StrToInt(E.CurrentData.CurrentCategory) == topic.Category.Id && E.ContainsTxt(topic) {
			if E.ContainsTxt(topic) {E.CurrentData.Topics = append(E.CurrentData.Topics, topic)}
		}
	}
}

/*
This function allows you to do a specific topic search with the str in parameters
*/
func (E *Engine) ContainsTxt(topic Topic) bool {
	if E.CurrentData.CurrentSearch == "" {
		return true
	} else { 
		return strings.Contains(strings.ToLower(topic.Title), strings.ToLower(E.CurrentData.CurrentSearch)) || strings.Contains(strings.ToLower(topic.Content), strings.ToLower(E.CurrentData.CurrentSearch))
	}
}


/*
This function allows you to retrieve the dislikes and likes from the database and define them on the Topic passed in prametters
It therefore takes as parameters a Topic and a string ("topicsLikes" or "topicsDislikes")
*/
func (E *Engine) SetLikedAndDisliked(table string, topic Topic) bool {
	data := E.QuerySQL("SELECT userID FROM " + table + " WHERE topicID = " + strconv.Itoa(topic.Id))
	var userID int
	for data.Next() {
		data.Scan(&userID)
		if (E.CurrentData.User.Id == userID) {
			return true
		}
	}
	return false
}
