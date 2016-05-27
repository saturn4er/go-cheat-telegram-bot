package tgbot

import (
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
	"log"
)

type Application struct {
	Token   string
	Clients map[int]*Client
	Debug   bool
}

func (a *Application) Run() {
	a.Clients = map[int]*Client{}
	bot, err := tgbotapi.NewBotAPI(a.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = a.Debug

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		var client = update.Message.From
		if _, ok := a.Clients[client.ID]; !ok {
			a.Clients[update.Message.From.ID] = NewClient(client, update.Message.Chat.ID, bot)
		}
		locClient := a.Clients[update.Message.From.ID]
		locClient.OnMessage(update.Message)
		if update.Message.Sticker != nil {
			fmt.Println("Received sticker with id: ", update.Message.Sticker.FileID)
		}
		if update.Message.Photo != nil {
			fmt.Println("Received photo with id: ", (*update.Message.Photo)[0].FileID)
		}
	}
}