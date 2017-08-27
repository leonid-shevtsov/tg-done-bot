package gtd_bot

import (
	"time"

	"github.com/go-pg/pg"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

type interactionState int

const (
	onboardingState = interactionState(iota)
	initialState
	isItActionableState
	whatIsTheGoalState
	whatIsTheNextActionState
	canYouDoItNowState
	doItNowState
	actionSuggestionState
	actionDoingState
	moveGoalForwardState
)

type interaction struct {
	repo    *repo
	bot     *telegram.BotAPI
	message *telegram.Message
	user    *User
}

func newInteraction(db *pg.DB, bot *telegram.BotAPI, message *telegram.Message) *interaction {
	return &interaction{repo: newRepo(db), bot: bot, message: message}
}

func (i *interaction) handleMessage() {
	defer i.repo.finalizeTransaction()

	i.user = i.repo.findUser(i.message.From.ID)

	// TODO perhaps show welcome text if it is a new user
	// TODO someday, check user payment status, or else

	i.user.LastMessageAt = time.Now()
	i.repo.update(i.user)

	// now dispatch based on user's state
	i.dispatchStateHandler()
}

func (i *interaction) dispatchStateHandler() {
	switch interactionState(i.user.State) {
	case onboardingState:
		i.handleOnboarding()
	case initialState:
		i.handleInitial()
	case isItActionableState:
		i.handleIsItActionable()
	case whatIsTheGoalState:
		i.handleWhatIsTheGoal()
	case whatIsTheNextActionState:
		i.handleWhatIsTheNextAction()
	case canYouDoItNowState:
		i.handleCanYouDoItNow()
	case doItNowState:
		i.handleDoItNow()
	case actionSuggestionState:
		i.handleActionSuggestion()
	case actionDoingState:
		i.handleActionDoing()
	case moveGoalForwardState:
		i.handleMoveGoalForward()
	default:
		panic("bad state")
	}
}

func (i *interaction) sendMessage(messageText string) {
	msg := telegram.NewMessage(int64(i.user.ID), messageText)
	i.bot.Send(msg)
}

func (i *interaction) sendPrompt(messageText string, keyboard [][]telegram.KeyboardButton) {
	msg := telegram.NewMessage(int64(i.user.ID), messageText)
	msg.ReplyMarkup = telegram.ReplyKeyboardMarkup{
		Keyboard:        keyboard,
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}
	i.bot.Send(msg)
}

func (i *interaction) sendUnclear() {
	i.sendMessage("Please pick one of the options.")
}

func (i *interaction) inboxCount() int {
	return i.repo.inboxCount(i.user.ID)
}

func (i *interaction) actionCount() int {
	return i.repo.actionCount(i.user.ID)
}
