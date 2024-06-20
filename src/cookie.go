package forum

import (
	"strconv"
	"time"
	"github.com/gofiber/fiber/v2"
)

func (E *Engine) GetCookieUser(c *fiber.Ctx) {
	if (c.Cookies("UserId", "false") == "false") {
		E.CurrentData.User = E.FindUserByID(E.StrToInt(c.Cookies("UserID")))
	}
}

func (E *Engine) SetCookieUser(userID int, c *fiber.Ctx) {
	cookieUser := new(fiber.Cookie)
	cookieUser.Name = "UserID"
	cookieUser.Value = strconv.Itoa(userID)
	cookieUser.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookieUser)
}


