package commands

import (
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
	"strings"
	"telegram-bot-framework"
)

type HelpCommand struct {
	tgbot.Command
}

func (hc *HelpCommand) Available(c *tgbot.Client) bool {
	return true
}
func (hc *HelpCommand) ExecuteForce(c *tgbot.Client) {
	sMsg := []string{}
	for _, command := range tgbot.GetCommands() {
		cmd := command.GetCommand()
		description := command.GetDescription()
		if command.Available(c) {
			if cmd != "" {
				sMsg = append(sMsg, fmt.Sprintf("%s - %s", cmd, description))
			}else if description != "" {
				sMsg = append(sMsg, description)
			}
		}
	}
	c.SendTextMessage(strings.Join(sMsg, "\n"))
}
func (hc *HelpCommand) Execute(m *tgbotapi.Message, c *tgbot.Client) bool {
	if m.Text == hc.GetCommand() {
		hc.ExecuteForce(c)
		return true
	}
	return false
}
func NewHelpCommand(command string, description string) *HelpCommand {
	result := new(HelpCommand)
	result.SetCommand(command)
	result.SetDescription(description)
	return result
}