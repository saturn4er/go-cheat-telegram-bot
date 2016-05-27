package tgbot

import (
	"gopkg.in/telegram-bot-api.v4"
)

var commands []Command

type Command interface {
	Available(c *Client) bool
	Execute(m *tgbotapi.Message, c *Client) bool
	ExecuteForce(c *Client)
	GetDescription() string
	GetCommand() string
}

func GetCommands() []Command {
	var result = []Command{}
	for _, c := range commands {
		result = append(result, c)
	}
	return result
}

func AddCommandExecutor(ce Command) {
	commands = append(commands, ce)
}