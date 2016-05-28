package commands

import (
	"gopkg.in/telegram-bot-api.v4"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/data"
)

type UnknownCommand struct {
	tgbot.Command
}

func (hc *UnknownCommand) Available(c *tgbot.Client) bool {
	return true
}
func (hc *UnknownCommand) ExecuteForce(c *tgbot.Client) {
	c.SendSticker(data.UnknownDataSticker)
	c.SendTextMessage("Я тебя не понимаю")
	c.ExecuteForceCommand(vHelpCommand)
}
func (nqc *UnknownCommand) Execute(m *tgbotapi.Message, c *tgbot.Client) bool {
	nqc.ExecuteForce(c)
	return false
}
func NewUnknownCommand() *UnknownCommand {
	result := new(UnknownCommand)
	return result
}