package gtd_bot

import (
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

const doingCommand = "Yes, I'll do this."
const skipCommand = "Skip this one for now."
const deferCommand = "Defer until later."

// const deferCommand = "Need a context."
const backToInboxCommand = "Done working for now."

var baseActionSuggestionKeyboard = [][]telegram.KeyboardButton{
	[]telegram.KeyboardButton{{Text: doingCommand}, {Text: skipCommand}, {Text: backToInboxCommand}},
}

func (i *interaction) actionSuggestionKeyboard() [][]telegram.KeyboardButton {
	baseKeyboard := baseActionSuggestionKeyboard
	// contexts := i.repo.contexts(i.user.ID)
	return baseKeyboard
}

func (i *interaction) gotoActionSuggestionState() {
	if actionToDo := i.repo.actionToDo(i.user.ID); actionToDo != nil {
		i.user.State = int(actionSuggestionState)
		i.user.CurrentActionID = actionToDo.ID
		i.user.CurrentGoalID = actionToDo.GoalID
		i.repo.update(i.user)

		i.sendMessage("I think you should do this now:")
		i.sendPrompt(actionToDo.Text, i.actionSuggestionKeyboard())
	} else {
		i.sendMessage("No more actions to do right now. Take a break?")
		i.gotoInitialState()
	}
}

func (i *interaction) handleActionSuggestion() {
	switch i.message.Text {
	case doingCommand:
		i.gotoActionDoingState()
	case skipCommand:
		i.skipAction()
		i.gotoActionSuggestionState()
	case backToInboxCommand:
		i.sendMessage("OK, back to collecting inbox.")
		i.gotoInitialState()
	default:
		i.sendUnclear()
		i.sendPrompt("Can you do this now?", i.actionSuggestionKeyboard())
	}
}

func (i *interaction) skipAction() {
	i.user.CurrentAction.ReviewedAt = time.Now()
	i.repo.update(i.user.CurrentAction)

	i.sendMessage("OK, skipping for now...")
}
