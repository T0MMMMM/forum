package forum

import (
	"strconv"
	"time"
	"github.com/gofiber/fiber/v2"
)

/*
Retrieving current user's cookies, this function is called on every page load
*/
func (E *Engine) GetCookieUser(c *fiber.Ctx) {
	if (c.Cookies("UserID") != "") {
		E.CurrentData.User = E.FindUserByID(E.StrToInt(c.Cookies("UserID")))
	} else {
		E.CurrentData.User = User{}
	}
}

/*
Retrieving cookies from the current category and current search so that the user can search for specific topics
*/
func (E *Engine) GetCookieFilters(c *fiber.Ctx) {
	E.GetCookieSearch(c)
	E.GetCookieCategory(c)
}

func (E *Engine) GetCookieSearch(c *fiber.Ctx) {
	E.CurrentData.CurrentSearch = c.Cookies("search")
}

func (E *Engine) GetCookieCategory(c *fiber.Ctx) {
	E.CurrentData.CurrentCategory = c.Cookies("category")
}

func (E *Engine) GetCookieTopic(c *fiber.Ctx) {
	if c.Cookies("topic") == "0" {
		E.CurrentData.Topic = Topic{Id: 0}
	} else if (c.Cookies("topic") != "") {
		E.CurrentData.Topic = E.FindTopicByID(E.StrToInt(c.Cookies("topic")))
		E.CurrentData.Topic.Liked = E.SetLikedAndDisliked("topicsLikes", E.CurrentData.Topic)
		E.CurrentData.Topic.Disliked = E.SetLikedAndDisliked("topicsDislikes", E.CurrentData.Topic)
	}
}

func (E *Engine) SetCookieUser(userID int, c *fiber.Ctx) {
	cookieUser := new(fiber.Cookie)
	cookieUser.Name = "UserID"
	cookieUser.Value = strconv.Itoa(userID)
	if (userID == 0) {
		cookieUser.Value = ""
	}
	cookieUser.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookieUser)
}

func (E *Engine) SetCookieSearch(search string, c *fiber.Ctx) {
	cookieSearch := new(fiber.Cookie)
	cookieSearch.Name = "search"
	cookieSearch.Value = search
	cookieSearch.Expires = time.Now().Add(5 * time.Hour)
	c.Cookie(cookieSearch)
}

func (E *Engine) SetCookieCategory(category string, c *fiber.Ctx) {
	cookieCategory := new(fiber.Cookie)
	cookieCategory.Name = "category"
	cookieCategory.Value = category
	cookieCategory.Expires = time.Now().Add(5 * time.Hour)
	c.Cookie(cookieCategory)
}

func (E *Engine) SetCookieTopic(idTopic int, c *fiber.Ctx) {
	cookieTopic := new(fiber.Cookie)
	cookieTopic.Name = "topic"
	cookieTopic.Value = strconv.Itoa(idTopic)
	cookieTopic.Expires = time.Now().Add(5 * time.Hour)
	c.Cookie(cookieTopic)
}


