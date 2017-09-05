package gtd_bot

import (
	"leonid.shevtsov.me/gtd_bot/i18n"

	"github.com/go-pg/pg"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

type interaction struct {
	bot     *telegram.BotAPI
	state   *state
	message *telegram.Message
	locale  *i18n.Locale
}

func handleMessage(bot *telegram.BotAPI, message *telegram.Message, db *pg.DB) {
	repo := newRepo(db)
	defer repo.finalizeTransaction()

	state := newState(repo, message.From.ID)
	interaction := interaction{state: state, bot: bot, message: message, locale: &i18n.En}
	interaction.state.setLastMessageNow()
	// TODO perhaps show welcome text if it is a new user
	// TODO someday, check user payment status, or else
	interaction.runQuestions()
}

const answerUnclear = "answer_unclear"

func (i *interaction) runQuestions() {
	question := questionMap[i.state.activeQuestion()]
	if question == nil {
		panic("unhandled question key")
		// TODO handle more gracefully?
	}

	nextQuestionKey := question.handleAnswer(i)

	if nextQuestionKey == answerUnclear {
		i.sendUnclear()
		return
	}

	nextQuestion := questionMap[nextQuestionKey]
	if nextQuestion == nil {
		panic("unhandled next question key")
	}

	i.state.setActiveQuestion(nextQuestionKey)
	nextQuestion.ask(i)
}

func (i *interaction) sendMessage(messageText string) {
	msg := telegram.NewMessage(int64(i.state.userID()), messageText)
	msg.ReplyMarkup = telegram.ReplyKeyboardRemove{RemoveKeyboard: true}
	i.bot.Send(msg)
}

func (i *interaction) sendPrompt(messageText string, keyboard [][]string) {
	var telegramKeyboard [][]telegram.KeyboardButton
	for _, row := range keyboard {
		var telegramRow []telegram.KeyboardButton
		for _, buttonText := range row {
			telegramRow = append(telegramRow, telegram.KeyboardButton{Text: buttonText})
		}
		telegramKeyboard = append(telegramKeyboard, telegramRow)
	}

	msg := telegram.NewMessage(int64(i.state.userID()), messageText)
	msg.ReplyMarkup = telegram.ReplyKeyboardMarkup{
		Keyboard:       telegramKeyboard,
		ResizeKeyboard: true,
	}
	i.bot.Send(msg)
}

func (i *interaction) sendUnclear() {
	msg := telegram.NewMessage(int64(i.state.userID()), i.locale.Messages.PickOneOfTheOptions)
	i.bot.Send(msg)
}
