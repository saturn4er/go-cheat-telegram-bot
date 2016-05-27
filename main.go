package main

import (
	"telegram-bot-framework"
	_ "github.com/saturn4er/go-cheat-telegram-bot/src/commands"
)

const BotToken = "BOT_TOKEN"

func main() {
	app := tgbot.Application{Token:BotToken}
	app.Run()
}
