package billet

import (
	"gopkg.in/telegram-bot-api.v4"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/data"
)

type CancelCommand struct {
	tgbot.Command
}

func (hc *CancelCommand) Available(c *tgbot.Client) bool {
	return c.GetState() == data.ClientSendingBilletId || c.GetState() == data.ClientSendingBilletQuestions
}
func (hc *CancelCommand) ExecuteForce(c *tgbot.Client) {
	c.SendTextMessage("Sorry we can't force run CancelCommand:(")
}
func (nqc *CancelCommand) Execute(m *tgbotapi.Message, c *tgbot.Client) bool {
	if !nqc.Available(c) {
		return false
	}
	if m.Text == nqc.GetCommand() {
		c.SendTextMessage("Ok :(")
		c.SetState(data.ClientIdle)
		SetUserBilletInfo(c, nil)
		return true
	}
	return false
}
func NewCancelCommand(command string, description string) *CancelCommand {
	result := new(CancelCommand)
	result.SetCommand(command)
	result.SetDescription(description)
	return result
}