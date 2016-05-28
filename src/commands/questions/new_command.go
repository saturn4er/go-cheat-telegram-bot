package questions

import (
	"gopkg.in/telegram-bot-api.v4"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/data"
	"github.com/saturn4er/go-cheat-telegram-bot/src/lib"
)

type NewQuestionCommand struct {
	tgbot.Command
}

func (hc *NewQuestionCommand) Available(c  *tgbot.Client) bool {
	return c.GetState() == data.ClientIdle
}
func (hc *NewQuestionCommand) ExecuteForce(c *tgbot.Client) {
	c.SetState(data.ClientSendingQuestionName)
	c.SendTextMessage("Укажите вопрос")
	newQuestion := lib.QuestionInfo{}
	newQuestion.Data = []lib.QuestionData{}
	SetUserQuestionInfo(c, &newQuestion)
}
func (nqc *NewQuestionCommand) Execute(m *tgbotapi.Message, c  *tgbot.Client) bool {
	if !nqc.Available(c) {
		return false
	}
	if m.Text == nqc.GetCommand() {
		nqc.ExecuteForce(c)
		return true
	}
	return false
}
func NewNewQuestionCommand(command string, description string) *NewQuestionCommand {
	result := new(NewQuestionCommand)
	result.SetCommand(command)
	result.SetDescription(description)
	return result
}