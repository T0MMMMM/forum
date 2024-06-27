package forum

import (
	"strconv"
)

func (E *Engine) CreateUserSearch(UserID int) UserSearch {
	user := E.FindUserByID(UserID)
	var (
		topics []Topic
		answersTopic []AnswerTopic
		idTopic int
		idAnswer int
		userID int
	)
	data := E.QuerySQL("SELECT id, userID FROM topics WHERE userID = " + strconv.Itoa(UserID))
	for data.Next() {
		data.Scan(&idTopic, &userID)
		topics = append(topics, E.FindTopicByID(idTopic))
	}
	data = E.QuerySQL("SELECT id, userID FROM answers WHERE userID = " + strconv.Itoa(UserID))
	for data.Next() {
		data.Scan(&idAnswer, &userID)
		answersTopic = append(answersTopic, AnswerTopic{Answer: E.FindAnswerByID(idAnswer), Topic: E.FindTopicByID(E.FindAnswerByID(idAnswer).TopicID)})
	}
	return UserSearch{Username: user.Username, Email: user.Email, CreatedAt: user.CreatedAt, ProfilePicture: user.ProfilePicture, AnswersTopic: answersTopic, Topics: topics}
}

func (E *Engine) FindTopicByID(TopicID int) Topic {
	data := E.QuerySQL("SELECT id, categoryID, userID, title, content, created_at, status, visible, like, dislike FROM topics WHERE id = " + strconv.Itoa(TopicID))
	var (
		id         int
		title      string
		categoryID int
		userID     int
		content    string
		created_at string
		status     string
		visible    bool
		like       int
		dislike    int
	)
	for data.Next() {data.Scan(&id, &categoryID, &userID, &title, &content, &created_at, &status, &visible, &like, &dislike)}

	return Topic{
		Id:        id,
		Title:     title,
		Content:   content,
		Category:  E.FindCategoryByID(categoryID),
		User:      E.FindUserByID(userID),
		CreatedAt: created_at,
		Status:    status,
		Visible:   visible,
		Like:      like,
		Dislike:   dislike,
		Answers:   E.FindAnswersByTopicID(TopicID),
	}
}

func (E *Engine) FindAnswerByID(AnswerID int) Answer {
	data := E.QuerySQL("SELECT id, userID, TopicID, content, created_at, status, visible, like, dislike FROM answers WHERE id = " + strconv.Itoa(AnswerID))
	var (
		id         int
		userID     int
		topicID    int
		content    string
		created_at string
		status     string
		visible    bool
		like       int
		dislike    int
	)
	for data.Next() {
		data.Scan(&id, &userID, &topicID, &content, &created_at, &status, &visible, &like, &dislike)
	}
	return Answer{Id: id, TopicID: topicID, Content: E.reversefilterMsg(content), CreatedAt: created_at, Status: status, Visible: visible, Like: like, Dislike: dislike}
}


func (E *Engine) FindCategoryByID(CategoryID int) Category {
	data := E.QuerySQL("SELECT id, name, description FROM Categories WHERE id = " + strconv.Itoa(CategoryID))
	var (
		id          int
		name        string
		description string
	)
	for data.Next() {
		data.Scan(&id, &name, &description)
	}
	return Category{
		Id:          id,
		Name:        name,
		Description: description,
	}
}


func (E *Engine) FindUserByID(UserID int) User {
	data := E.QuerySQL("SELECT id, email, username, password, created_at, profile_picture FROM users WHERE id = " + strconv.Itoa(UserID))
	var (
		id         int
		email      string
		username   string
		password   string
		created_at string
		profilePicture string
	)
	for data.Next() {
		data.Scan(&id, &email, &username, &password, &created_at, &profilePicture)
	}
	return User{
		Id:        id,
		Username:  username,
		Email:     email,
		CreatedAt: created_at,
		ProfilePicture: profilePicture,
	}
}

func (E *Engine) FindAnswersByTopicID(TopicID int) []Answer {
	data := E.QuerySQL("SELECT id, userID, content, created_at, status, visible, like, dislike FROM answers WHERE topicID = " + strconv.Itoa(TopicID))
	var (
		answers    []Answer
		id         int
		userID     int
		content    string
		created_at string
		status     string
		visible    bool
		like       int
		dislike    int
	)
	for data.Next() {
		data.Scan(&id, &userID, &content, &created_at, &status, &visible, &like, &dislike)
		answers = append(answers, Answer{Id: id, TopicID: TopicID, User: E.FindUserByID(userID), Content: E.reversefilterMsg(content), CreatedAt: created_at, Status: status, Visible: visible, Like: like, Dislike: dislike})
	}
	return answers
}

func (E *Engine) FindMessageByID(msgId int) Message {
	data := E.QuerySQL("SELECT id, senderID, recipientID, content, visible, created_at FROM messages WHERE id = " + strconv.Itoa(msgId))
	var (
		id         int
		senderID     int
		recipientID    int
		content 	string
		visible     bool
		created_at string
		
	)
	for data.Next() {
		data.Scan(&id, &senderID, &recipientID, &content, &visible, &created_at)
	}
	return Message{Id: id, Sender: E.FindUserByID(senderID), Recipient: E.FindUserByID(recipientID), Content: content, Visible: visible, CreatedAt: created_at}
}

