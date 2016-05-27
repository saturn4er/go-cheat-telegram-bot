package data

import "telegram-bot-framework"

const (
	ClientIdle tgbot.ClientState = iota
	ClientSendingQuestionName
	ClientSendingQuestionData
)