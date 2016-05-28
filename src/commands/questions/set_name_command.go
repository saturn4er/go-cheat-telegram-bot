package questions

import (
	"gopkg.in/telegram-bot-api.v4"
	"crypto/md5"
	"time"
	"encoding/hex"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/data"
)

type SetQuestionNameCommand struct {
	tgbot.Command
}

func (this *SetQuestionNameCommand) Available(c  *tgbot.Client) bool {
	return c.GetState() == data.ClientSendingQuestionName
}
func (hc *SetQuestionNameCommand) ExecuteForce(c *tgbot.Client) {
	c.SendTextMessage("Sorry we can't force run SetQuestionNameCommand:(")
}
func (this *SetQuestionNameCommand) Execute(m *tgbotapi.Message, c  *tgbot.Client) bool {
	if !this.Available(c) {
		return false
	}
	if m.Text == "" {
		c.SendTextMessage("Укажите правильное название")
		return true
	}
	c.SetState(data.ClientSendingQuestionData)
	question := GetUserQuestionInfo(c)
	question.Name = m.Text
	nameMD5 := md5.Sum([]byte(question.Name + time.Now().String()))
	question.Hash = hex.EncodeToString(nameMD5[:])
	SetUserQuestionInfo(c, question)
	c.SendTextMessage("Теперь нужно указать данные, которые я тебе отправлю, когда ты будешь списывать;) Сейчас доступны только: текст, картинка, документ")
	return true
}
func NewSetQuestionNameCommand(description string) *SetQuestionNameCommand {
	result := new(SetQuestionNameCommand)
	result.SetDescription(description)
	return result
}