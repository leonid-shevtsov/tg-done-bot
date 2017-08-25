package gtd_bot

import (
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

const doneCommand = "Done!"
const doItLaterCommand = "I'll do it later"

var doItNowKeyboard = [][]telegram.KeyboardButton{
	[]telegram.KeyboardButton{{Text: doneCommand}, {Text: doItLaterCommand}},
}

func (i *interaction) gotoDoItNow() {
	i.user.State = int(doItNowState)
	i.repo.update(i.user)
	i.sendPrompt("Do it! 2 minutes and counting.", doItNowKeyboard)
	// TODO - notify with timer
	// go i.setDoItNowTimer(i.user.ID, i.user.CurrentActionID)
}

func (i *interaction) handleDoItNow() {
	switch i.message.Text {
	case doneCommand:
		i.markActionCompleted()
		i.gotoProcessInbox()
	case doItLaterCommand:
		i.gotoProcessInbox()
	default:
		i.sendUnclear()
		i.gotoDoItNow()
	}
}

func (i *interaction) markActionCompleted() {
	i.user.CurrentAction.CompletedAt = time.Now()
	i.repo.update(i.user.CurrentAction)
}
