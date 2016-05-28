package billet

import (
	"gopkg.in/telegram-bot-api.v4"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/data"
	"fmt"
)

type EndCommand struct {
	tgbot.Command
}

func (this *EndCommand) Available(c  *tgbot.Client) bool {
	return c.GetState() == data.ClientSendingBilletQuestions
}
func (hc *EndCommand) ExecuteForce(c *tgbot.Client) {
	c.SendTextMessage("Sorry we can't force run EndCommand:(")
}
func (this *EndCommand) Execute(m *tgbotapi.Message, c  *tgbot.Client) bool {
	if !this.Available(c) {
		return false
	}
	if m.Text == this.GetCommand() {
		info := GetUserBilletInfo(c)
		err := info.Save()
		if err != nil {
			fmt.Println(err)
			c.SendSticker(data.ErrorSticker)
			c.SendTextMessage("Что-то пошло не так. Пожешь попытаться добавить снова(/end) или отменить (/cancel)")
			return true
		}
		c.SetState(data.ClientIdle)
		c.SendSticker(data.YeahSticker)
		c.SendTextMessage("Добавлено")
		return true
	}
	return false
}
func NewEndCommand(command, description string) *EndCommand {
	result := new(EndCommand)
	result.SetCommand(command)
	result.SetDescription(description)
	return result
}