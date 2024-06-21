package forum

import (
	"strconv"
)

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
	data := E.QuerySQL("SELECT id, email, username, password, created_at FROM users WHERE id = " + strconv.Itoa(UserID))
	var (
		id         int
		email      string
		username   string
		password   string
		created_at string
	)
	for data.Next() {
		data.Scan(&id, &email, &username, &password, &created_at)
	}
	return User{
		Id:        id,
		Username:  username,
		Email:     email,
		CreatedAt: created_at,
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
		answers = append(answers, Answer{Id: id, User: E.FindUserByID(userID), Content: E.reversefilterMsg(content), CreatedAt: created_at, Status: status, Visible: visible, Like: like, Dislike: dislike})
	}
	return answers
}

