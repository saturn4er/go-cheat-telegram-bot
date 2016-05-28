package tgbot

import (
	"gopkg.in/telegram-bot-api.v4"
)

var commands []commander

type Command struct {
	command     string
	description string
}
type commander interface {
	Available(c *Client) bool
	Execute(m *tgbotapi.Message, c *Client) bool
	ExecuteForce(c *Client)
	GetDescription() string
	GetCommand() string
	SetDescription(string)
	SetCommand(string)
}

func (cmd *Command) Available(c *Client) bool {
	return false
}
func (cmd *Command) Execute(m *tgbotapi.Message, c *Client) bool {
	return false
}
func (cmd *Command) ExecuteForce(c *Client) {

}
func (cmd *Command) GetDescription() string {
	return cmd.description
}
func (cmd *Command) GetCommand() string {
	return cmd.command
}
func (cmd *Command) SetDescription(description string) {
	cmd.description = description
}
func (cmd *Command) SetCommand(command string) {
	cmd.command = command
}

func GetCommands() []commander {
	var result = []commander{}
	for _, c := range commands {
		result = append(result, c)
	}
	return result
}

func AddCommandExecutor(ce commander) {
	commands = append(commands, ce)
}
