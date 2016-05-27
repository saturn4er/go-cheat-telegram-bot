package questions

import (
	"gopkg.in/telegram-bot-api.v4"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/data"
)

type CancelNewQuestionCommand struct {
	command     string
	description string
}

func (hc *CancelNewQuestionCommand) Available(c *tgbot.Client) bool {
	return c.GetState() == data.ClientSendingQuestionName || c.GetState() == data.ClientSendingQuestionData
}
func (hc *CancelNewQuestionCommand) ExecuteForce(c *tgbot.Client) {
	c.SendTextMessage("Sorry we can't force run CancelNewQuestionCommand:(")
}
func (nqc *CancelNewQuestionCommand) Execute(m *tgbotapi.Message, c *tgbot.Client) bool {
	if !nqc.Available(c) {
		return false
	}
	if m.Text == nqc.GetCommand() {
		c.SendTextMessage("Ok :(")
		c.SetState(data.ClientIdle)
		c.SetData("question", nil)
		return true
	}
	return false
}
func (hc *CancelNewQuestionCommand) GetCommand() string {
	return hc.command
}
func (hc *CancelNewQuestionCommand) GetDescription() string {
	return hc.description
}
func NewCancelNewQuestionCommand(command string, description string) *CancelNewQuestionCommand {
	result := new(CancelNewQuestionCommand)
	result.command = command
	result.description = description
	return result
}