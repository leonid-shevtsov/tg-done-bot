package gtd_bot

import (
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

const yesCommand = "Yes"
const noCommand = "No - trash it"
const abortCommand = "Let's do this later"

var isItActionableKeyboard = [][]telegram.KeyboardButton{
	[]telegram.KeyboardButton{{Text: yesCommand}, {Text: noCommand}, {Text: abortCommand}},
}

func (i *interaction) gotoProcessInbox() bool {
	return i.gotoIsItActionable()
}

func (i *interaction) gotoIsItActionable() bool {
	if inboxItemToProcess := i.repo.inboxItemToProcess(i.user.ID); inboxItemToProcess != nil {
		i.user.State = int(isItActionableState)
		i.user.CurrentInboxItemID = inboxItemToProcess.ID
		i.repo.update(i.user)

		i.sendMessage("Processing inbox item:")
		i.sendMessage(inboxItemToProcess.Text)
		i.sendPrompt("Is it actionable?", isItActionableKeyboard)
		return true
	} else {
		i.sendMessage("Inbox zero!")
		return false
	}
}

func (i *interaction) handleIsItActionable() {
	switch i.message.Text {
	case yesCommand:
		i.gotoWhatIsTheGoal()
	case noCommand:
		i.trashCurrentInboxItem()
	case abortCommand:
		i.abortProcessing()
	default:
		i.sendUnclear()
		i.sendPrompt("Is it actionable?", isItActionableKeyboard)
	}
}

func (i *interaction) trashCurrentInboxItem() {
	i.user.CurrentInboxItem.ProcessedAt = time.Now()
	i.repo.update(i.user.CurrentInboxItem)

	i.sendMessage("Trashed! Moving on.")
	if !i.gotoProcessInbox() {
		i.gotoInitialState()
	}
}
