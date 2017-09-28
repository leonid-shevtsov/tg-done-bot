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
