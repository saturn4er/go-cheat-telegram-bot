package commands

import (
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
	"encoding/json"
	"strings"
	"telegram-bot-framework"
)

type DebugCommand struct {
	tgbot.Command
}

func (hc *DebugCommand) Available(c *tgbot.Client) bool {
	return true
}
func (hc *DebugCommand) ExecuteForce(c *tgbot.Client) {
	sMsg := []string{fmt.Sprintf("State: %d", c.GetState())}
	for key, value := range c.GetAllData() {
		line := key + " - "
		val, err := json.Marshal(value)
		if err == nil {
			line += string(val)
		}else {
			line += fmt.Sprintf("%+v", value)
		}
		sMsg = append(sMsg, line)
	}
	c.SendTextMessage(strings.Join(sMsg, "\n"))
}
func (nqc *DebugCommand) Execute(m *tgbotapi.Message, c *tgbot.Client) bool {
	if m.Text == nqc.GetCommand() {
		nqc.ExecuteForce(c)
		return true
	}
	return false
}
func NewDebugCommand(command string, description string) *DebugCommand {
	result := new(DebugCommand)
	result.SetCommand(command)
	result.SetDescription(description)
	return result
}