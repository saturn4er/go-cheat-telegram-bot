package commands

import (
	"github.com/saturn4er/go-cheat-telegram-bot/src/commands/questions"
	"telegram-bot-framework"
)

var (
	vHelpCommand = NewHelpCommand("/help", "Помощь")
	vDebugCommand = NewDebugCommand("/debug", "Дебаг")
	vQuestionsListCommand = questions.NewQuestionsListCommand("/qlist", "Список вопросов")
	vCancelNewQuestionCommand = questions.NewCancelNewQuestionCommand("/cancel", "Отмена")
	vNewQuestionCommand = questions.NewNewQuestionCommand("/newquestion", "Добавить вопрос")
	vEndCommand = questions.NewEndCommand("/end", "Введите название вопроса")
	vAddDataCommand = questions.NewAddDataCommand("Отправьте данные")
	vSetQuestionNameCommand = questions.NewSetQuestionNameCommand("Введите название вопроса")
	vUnknownCommand = NewUnknownCommand()
)

func init() {
	tgbot.AddCommandExecutor(vHelpCommand)
	tgbot.AddCommandExecutor(vDebugCommand)
	tgbot.AddCommandExecutor(vQuestionsListCommand)
	tgbot.AddCommandExecutor(vCancelNewQuestionCommand)
	tgbot.AddCommandExecutor(vNewQuestionCommand)
	tgbot.AddCommandExecutor(vEndCommand)
	tgbot.AddCommandExecutor(vAddDataCommand)
	tgbot.AddCommandExecutor(vSetQuestionNameCommand)
	tgbot.AddCommandExecutor(vUnknownCommand)

}