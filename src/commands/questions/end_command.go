package questions

import (
	"gopkg.in/telegram-bot-api.v4"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/data"
)

type EndCommand struct {
	command     string
	description string
}

func (this *EndCommand) Available(c  *tgbot.Client) bool {
	return c.GetState() == data.ClientSendingQuestionData
}
func (hc *EndCommand) ExecuteForce(c *tgbot.Client) {
	c.SendTextMessage("Sorry we can't force run EndCommand:(")
}
func (this *EndCommand) Execute(m *tgbotapi.Message, c  *tgbot.Client) bool {
	if !this.Available(c) {
		return false
	}
	if m.Text == this.GetCommand() {
		c.SetState(data.ClientIdle)
		c.SendSticker(data.YeahSticker)
		c.SendTextMessage("Добавлено")
		return true
	}
	return false
}
func (this *EndCommand) GetCommand() string {
	return this.command
}
func (this *EndCommand) GetDescription() string {
	return this.description
}
func NewEndCommand(command, description string) *EndCommand {
	result := new(EndCommand)
	result.command = command
	result.description = description
	return result
}