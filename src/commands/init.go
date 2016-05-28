package commands

import (
	"github.com/saturn4er/go-cheat-telegram-bot/src/commands/questions"
	"telegram-bot-framework"
	"github.com/saturn4er/go-cheat-telegram-bot/src/commands/billet"
)

var (
	vHelpCommand = NewHelpCommand("/help", "Помощь")
	vDebugCommand = NewDebugCommand("/debug", "Дебаг")

	vQuestionsListCommand = questions.NewQuestionsListCommand("/qlist", "Список вопросов")
	vCancelNewQuestionCommand = questions.NewCancelNewQuestionCommand("/cancel", "Отмена")
	vNewQuestionCommand = questions.NewNewQuestionCommand("/newquestion", "Добавить вопрос")
	vEndCommand = questions.NewEndCommand("/end", "Завершить")
	vAddDataCommand = questions.NewAddDataCommand("Отправьте данные")
	vSetQuestionNameCommand = questions.NewSetQuestionNameCommand("Введите название вопроса")

	vBilletListCommand = billet.NewListCommand("/blist", "Список билетов")
	vNewBilletCommand = billet.NewNewBilletCommand("/newbillet", "Добавить билет")
	vSetBilletIdCommand = billet.NewSetIdCommand("Отправьте идентификатор билета")
	vCancelBilletCommand = billet.NewCancelCommand("/cancel", "Отмена")
	vAddQuestionNewBilletCommand = billet.NewAddQuestionNewBilletCommand("Отправьте хеш-сумму вопроса")
	vEndBilletCommand = billet.NewEndCommand("/end", "Завершить")

	vUnknownCommand = NewUnknownCommand()
)

func init() {
	tgbot.AddCommandExecutor(vHelpCommand)
	tgbot.AddCommandExecutor(vDebugCommand)

	// questions
	tgbot.AddCommandExecutor(vQuestionsListCommand)
	tgbot.AddCommandExecutor(vCancelNewQuestionCommand)
	tgbot.AddCommandExecutor(vNewQuestionCommand)
	tgbot.AddCommandExecutor(vEndCommand)
	tgbot.AddCommandExecutor(vAddDataCommand)
	tgbot.AddCommandExecutor(vSetQuestionNameCommand)

	// billets
	tgbot.AddCommandExecutor(vBilletListCommand)
	tgbot.AddCommandExecutor(vNewBilletCommand)
	tgbot.AddCommandExecutor(vSetBilletIdCommand)
	tgbot.AddCommandExecutor(vCancelBilletCommand)
	tgbot.AddCommandExecutor(vEndBilletCommand)
	tgbot.AddCommandExecutor(vAddQuestionNewBilletCommand)

	tgbot.AddCommandExecutor(vUnknownCommand)

}