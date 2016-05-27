package questions

import (
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/lib"
)

func GetUserQuestionInfo(c *tgbot.Client) *lib.QuestionInfo {
	data := c.GetData("question")
	if val, ok := data.(*lib.QuestionInfo); ok {
		return val
	}
	return nil
}
func SetUserQuestionInfo(c *tgbot.Client, data *lib.QuestionInfo) {
	c.SetData("question", data)
}