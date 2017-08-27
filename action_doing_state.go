package gtd_bot

import telegram "github.com/go-telegram-bot-api/telegram-bot-api"

var actionDoingKeyboard = [][]telegram.KeyboardButton{
	[]telegram.KeyboardButton{{Text: doneCommand}, {Text: doItLaterCommand}},
}

func (i *interaction) gotoActionDoingState() {
	i.user.State = int(actionDoingState)
	i.repo.update(i.user)

	i.sendPrompt("Great! Waiting for you to finish.", actionDoingKeyboard)
}

func (i *interaction) handleActionDoing() {
	switch i.message.Text {
	case doneCommand:
		i.markActionCompleted()
		i.sendMessage("Awesome!")
		i.gotoMoveGoalForward()
	case doItLaterCommand:
		i.skipAction()
		i.gotoActionSuggestionState()
	default:
		i.sendUnclear()
		i.sendPrompt("Waiting for you to finish.", actionDoingKeyboard)
	}
}
