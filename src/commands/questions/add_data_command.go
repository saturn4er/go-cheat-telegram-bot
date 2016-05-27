package questions

import (
	"gopkg.in/telegram-bot-api.v4"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/lib"
	"github.com/saturn4er/go-cheat-telegram-bot/src/data"
)

type AddDataCommand struct {
	command     string
	description string
}

func (this *AddDataCommand) Available(c  *tgbot.Client) bool {
	return c.GetState() == data.ClientSendingQuestionData
}
func (hc *AddDataCommand) ExecuteForce(c *tgbot.Client) {
	c.SendTextMessage("Sorry we can't force run AddDataCommand:(")
}
func (this *AddDataCommand) Execute(m *tgbotapi.Message, c  *tgbot.Client) bool {
	if !this.Available(c) {
		return false
	}

	question := GetUserQuestionInfo(c)
	if m.Text != "" {
		question.Data = append(question.Data, lib.QuestionData{Data:m.Text, Type:lib.QuestionDataText})
	} else if m.Document != nil {
		question.Data = append(question.Data, lib.QuestionData{Data:m.Document.FileID, Type:lib.QuestionDataDocument})
	} else if m.Photo != nil && len(*m.Photo) > 0 {
		question.Data = append(question.Data, lib.QuestionData{Data:(*m.Photo)[0].FileID, Type:lib.QuestionDataPhoto})
	} else {
		c.SendPhoto(data.YouWillDiePhoto)
		return true
	}
	SetUserQuestionInfo(c, question)
	msg := tgbotapi.NewMessage(c.GetChatID(), "Можешь отправить ещё что-то.\nЕсли больше ничего нет - отправь /end")
	c.Send(msg)
	return true
}
func (this *AddDataCommand) GetCommand() string {
	return this.command
}
func (this *AddDataCommand) GetDescription() string {
	return this.description
}
func NewAddDataCommand(description string) *AddDataCommand {
	result := new(AddDataCommand)
	result.command = ""
	result.description = description
	return result
}