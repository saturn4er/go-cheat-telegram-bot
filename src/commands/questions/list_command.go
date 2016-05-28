package questions

import (
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
	"strings"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/data"
	"github.com/saturn4er/go-cheat-telegram-bot/src/lib"
)

type QuestionsListCommand struct {
	tgbot.Command
}

func (qlc *QuestionsListCommand) Available(c  *tgbot.Client) bool {
	return c.GetState() == data.ClientIdle
}
func (hc *QuestionsListCommand) ExecuteForce(c *tgbot.Client) {
	questions, err := lib.GetQuestionsList()
	if err != nil {
		c.SendTextMessage("Не пошлоо")
		return
	}
	sMsg := []string{}
	for _, q := range questions {
		sMsg = append(sMsg, fmt.Sprintf("/%s - %s", q.Hash, q.Name))
	}
	c.SendTextMessage(strings.Join(sMsg, "\n"))
}
func (qlc *QuestionsListCommand) Execute(m *tgbotapi.Message, c  *tgbot.Client) bool {
	if !qlc.Available(c) {
		return false
	}
	if m.Text == qlc.GetCommand() {
		qlc.ExecuteForce(c)
		return true
	}
	return false
}
func NewQuestionsListCommand(command string, description string) *QuestionsListCommand {
	result := new(QuestionsListCommand)
	result.SetCommand(command)
	result.SetDescription(description)
	return result
}