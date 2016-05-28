package billet

import (
	"gopkg.in/telegram-bot-api.v4"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/data"
	"github.com/saturn4er/go-cheat-telegram-bot/src/lib"
)

type NewBilletCommand struct {
	tgbot.Command
}

func (hc *NewBilletCommand) Available(c  *tgbot.Client) bool {
	return c.GetState() == data.ClientIdle
}
func (hc *NewBilletCommand) ExecuteForce(c *tgbot.Client) {
	c.SetState(data.ClientSendingBilletId)
	c.SendTextMessage("Укажите идентификатор билета(номер или название)")
	newBillet := lib.BilletInfo{}
	SetUserBilletInfo(c, &newBillet)
}
func (nqc *NewBilletCommand) Execute(m *tgbotapi.Message, c  *tgbot.Client) bool {
	if !nqc.Available(c) {
		return false
	}
	if m.Text == nqc.GetCommand() {
		nqc.ExecuteForce(c)
		return true
	}
	return false
}
func NewNewBilletCommand(command string, description string) *NewBilletCommand {
	result := new(NewBilletCommand)
	result.SetCommand(command)
	result.SetDescription(description)
	return result
}