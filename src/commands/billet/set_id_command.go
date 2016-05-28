package billet

import (
	"gopkg.in/telegram-bot-api.v4"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/data"
	"crypto/md5"
	"time"
	"encoding/hex"
)

type SetIdCommand struct {
	tgbot.Command
}

func (hc *SetIdCommand) Available(c  *tgbot.Client) bool {
	return c.GetState() == data.ClientSendingBilletId
}
func (hc *SetIdCommand) ExecuteForce(c *tgbot.Client) {
	c.SendTextMessage("Cannot force execute SetIdCommand:(")
}
func (nqc *SetIdCommand) Execute(m *tgbotapi.Message, c  *tgbot.Client) bool {
	if !nqc.Available(c) {
		return false
	}
	c.SetState(data.ClientSendingBilletQuestions)
	c.SendTextMessage("А теперь скинь хеш суммы вопросов, которые в этом билете")
	billet := GetUserBilletInfo(c)
	billet.Id = m.Text
	nameMD5 := md5.Sum([]byte(m.Text + time.Now().String()))
	billet.Hash = hex.EncodeToString(nameMD5[:])
	SetUserBilletInfo(c, billet)
	return true
}
func NewSetIdCommand(description string) *SetIdCommand {
	result := new(SetIdCommand)
	result.SetDescription(description)
	return result
}