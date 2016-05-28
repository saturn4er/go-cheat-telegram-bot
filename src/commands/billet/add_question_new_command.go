package billet

import (
	"gopkg.in/telegram-bot-api.v4"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/data"
	"github.com/saturn4er/go-cheat-telegram-bot/src/lib"
)

type AddQuestionNewBilletCommand struct {
	tgbot.Command
}

func (hc *AddQuestionNewBilletCommand) Available(c  *tgbot.Client) bool {
	return c.GetState() == data.ClientSendingBilletQuestions
}
func (hc *AddQuestionNewBilletCommand) ExecuteForce(c *tgbot.Client) {
	c.SendTextMessage("Cannot force execute SetIdCommand:(")
}
func (nqc *AddQuestionNewBilletCommand) Execute(m *tgbotapi.Message, c  *tgbot.Client) bool {
	if !nqc.Available(c) {
		return false
	}
	if m.Text == ""{
		c.SendTextMessage(`Отправь мне хеш сумму!(Тот текст, что написан после слеша когда ты запрашиваешь вопрос после "question_").
Если больше нечего отправить - напиши /end
Если хочешь отменить создание билета-шли /cancel`)
		return true
	}
	billets, err := lib.GetQuestionsList()
	if err != nil {
		c.SendSticker(data.ErrorSticker)
		return true
	}
	var found bool
	for _, b := range billets {
		if b.Hash == m.Text {
			found = true
			break
		}
	}
	if !found {
		c.SendSticker(data.UnknownDataSticker)
		c.SendTextMessage("Я такого вопроса найти не могу:(")
		return true
	}
	c.SendTextMessage("Ещё что-то? Если нет - отправь /end\nЕсли хочешь отменить создание билета-шли /cancel")
	billet := GetUserBilletInfo(c)
	billet.Questions = append(billet.Questions, m.Text)
	SetUserBilletInfo(c, billet)
	return true
}
func NewAddQuestionNewBilletCommand(description string) *AddQuestionNewBilletCommand {
	result := new(AddQuestionNewBilletCommand)
	result.SetDescription(description)
	return result
}