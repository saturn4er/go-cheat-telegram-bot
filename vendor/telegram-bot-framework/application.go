package tgbot

import (
	"gopkg.in/telegram-bot-api.v4"
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

	log.Printf("Authorized on account %s", bot.Self.UserName)
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
	}
}