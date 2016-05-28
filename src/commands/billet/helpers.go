package billet

import (
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/lib"
)

func GetUserBilletInfo(c *tgbot.Client) *lib.BilletInfo {
	data := c.GetData("question")
	if val, ok := data.(*lib.BilletInfo); ok {
		return val
	}
	return nil
}
func SetUserBilletInfo(c *tgbot.Client, data *lib.BilletInfo) {
	c.SetData("question", data)
}