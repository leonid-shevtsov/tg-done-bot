package tg_done_bot

import (
	"strings"

	"github.com/leonid-shevtsov/tg-done-bot/i18n"

	"github.com/go-pg/pg"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

type interaction struct {
	bot     *telegram.BotAPI
	state   *state
	message *telegram.Message
	locale  *i18n.Locale
}

func initiateInteraction(bot *telegram.BotAPI, repo *repo, userID int) *interaction {
	state := newState(repo, userID)
	return &interaction{state: state, bot: bot, message: nil, locale: &i18n.En}
}

func handleMessage(bot *telegram.BotAPI, message *telegram.Message, db *pg.DB) {
	repo := newRepo(db)
	defer repo.finalizeTransaction()

	state := newState(repo, message.From.ID)
	interaction := interaction{state: state, bot: bot, message: message, locale: &i18n.En}
	interaction.state.setLastMessageNow()
	// TODO someday, check user payment status, or else
	if interaction.message.Text[0] == '/' {
		interaction.runCommands()
	} else {
		interaction.runQuestions()
	}
}

func (i *interaction) runCommands() {
	messageParts := strings.Split(i.message.Text, " ")
	command := messageParts[0]
	arguments := messageParts[1:]
	switch command {
	case "/abort":
		commandAbort(i, arguments)
	case "/inbox":
		commandInbox(i, arguments)
	case "/goals":
		commandGoals(i, arguments)
	case "/actions":
		commandActions(i, arguments)
	case "/waiting":
		commandWaiting(i, arguments)
	case "/contexts":
		commandContexts(i, arguments)
	default:
		i.reply().text(i.locale.Slash.CommandUnknown).send()
	}
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

	i.state.setActiveQuestion(nextQuestionKey)
	i.askActiveQuestion()
}

func (i *interaction) askActiveQuestion() {
	activeQuestion := questionMap[i.state.activeQuestion()]
	if activeQuestion == nil {
		panic("unhandled next question key")
	}
	activeQuestion.ask(i)
}
