package tgbot

import (
	"gopkg.in/telegram-bot-api.v4"
)

type ClientState int
type Client struct {
	user   *tgbotapi.User
	state  ClientState
	bot    *tgbotapi.BotAPI
	chatID int64
	data   map[string]interface{}
}

func NewClient(client *tgbotapi.User, chatId int64, bot *tgbotapi.BotAPI) *Client {
	c := new(Client)
	c.user = client
	c.chatID = chatId
	c.bot = bot
	c.data = map[string]interface{}{}
	return c
}
func (c *Client) Send(message tgbotapi.Chattable) {
	c.bot.Send(message)
}
func (c *Client) SendTextMessage(text string) {
	msg := tgbotapi.NewMessage(c.chatID, text)
	c.Send(msg)
}
func (c *Client) SendSticker(fileId string) {
	msg := tgbotapi.NewStickerShare(c.chatID, fileId)
	c.Send(msg)
}
func (c *Client) SendPhoto(fileId string) {
	msg := tgbotapi.NewPhotoShare(c.chatID, fileId)
	c.Send(msg)
}
func (c *Client) SendSequence(s []tgbotapi.Chattable) {
	for _, message := range s {
		c.Send(message)
	}
}
func (c *Client) GetChatID() int64 {
	return c.chatID
}
func (c *Client) GetState() ClientState {
	return c.state
}
func (c *Client) SetState(state ClientState) {
	c.state = state
}
func (c *Client) GetTGClient() *tgbotapi.User {
	return c.user
}
func (c *Client) GetData(param string) interface{} {
	return c.data[param]
}
func (c *Client) GetAllData() map[string]interface{} {
	return c.data
}
func (c *Client) SetData(param string, data interface{}) {
	c.data[param] = data
}
func (c *Client) OnMessage(message *tgbotapi.Message) {
	for _, command := range GetCommands() {
		if c.ExecuteCommand(message, command) {
			break
		}
	}
	return

}
func (c *Client) ExecuteCommand(message *tgbotapi.Message, command Command) bool {
	return command.Available(c) && command.Execute(message, c)
}
func (c *Client) ExecuteForceCommand(command Command) {
	command.ExecuteForce(c)
}