package billet

import (
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
	"strings"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/data"
	"github.com/saturn4er/go-cheat-telegram-bot/src/lib"
)

type ListCommand struct {
	tgbot.Command
}

func (qlc *ListCommand) Available(c  *tgbot.Client) bool {
	return c.GetState() == data.ClientIdle
}
func (hc *ListCommand) ExecuteForce(c *tgbot.Client) {
	questions, err := lib.GetBilletsList()
	if err != nil {
		c.SendTextMessage("Не пошлоо")
		return
	}
	sMsg := []string{}
	for _, q := range questions {
		sMsg = append(sMsg, fmt.Sprintf("/billet_%s - %s", q.Hash, q.Id))
	}
	c.SendTextMessage(strings.Join(sMsg, "\n"))
}
func (qlc *ListCommand) Execute(m *tgbotapi.Message, c  *tgbot.Client) bool {
	if !qlc.Available(c) {
		return false
	}
	if m.Text == qlc.GetCommand() {
		qlc.ExecuteForce(c)
		return true
	}
	return false
}
func NewListCommand(command string, description string) *ListCommand {
	result := new(ListCommand)
	result.SetCommand(command)
	result.SetDescription(description)
	return result
}