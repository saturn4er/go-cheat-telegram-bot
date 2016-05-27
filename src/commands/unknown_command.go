package commands

import (
	"gopkg.in/telegram-bot-api.v4"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/data"
)

type UnknownCommand struct {
	command     string
	description string
}

func (hc *UnknownCommand) Available(c *tgbot.Client) bool {
	return true
}
func (hc *UnknownCommand) ExecuteForce(c *tgbot.Client) {
	c.SendSticker(data.UnknownDataSticker)
	c.ExecuteForceCommand(vHelpCommand)
}
func (nqc *UnknownCommand) Execute(m *tgbotapi.Message, c *tgbot.Client) bool {
	nqc.ExecuteForce(c)
	return false
}
func (hc *UnknownCommand) GetCommand() string {
	return hc.command
}
func (hc *UnknownCommand) GetDescription() string {
	return hc.description
}
func NewUnknownCommand() *UnknownCommand {
	result := new(UnknownCommand)
	result.command = ""
	result.description = ""
	return result
}