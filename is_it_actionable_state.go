package gtd_bot

import (
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

const yesCommand = "Yes"
const noTrashItCommand = "No - trash it"
const abortCommand = "Let's do this later"

var isItActionableKeyboard = [][]telegram.KeyboardButton{
	[]telegram.KeyboardButton{{Text: yesCommand}, {Text: noTrashItCommand}, {Text: abortCommand}},
}

func (i *interaction) gotoProcessInbox() {
	if inboxItemToProcess := i.repo.inboxItemToProcess(i.user.ID); inboxItemToProcess != nil {
		i.user.State = int(isItActionableState)
		i.user.CurrentInboxItemID = inboxItemToProcess.ID
		i.user.CurrentGoalID = 0

		i.repo.update(i.user)

		i.sendMessage("Processing inbox item:")
		i.sendMessage(inboxItemToProcess.Text)
		i.sendPrompt("Is it actionable?", isItActionableKeyboard)
	} else {
		i.sendMessage("Inbox zero!")
		i.gotoInitialState()
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
	i.gotoProcessInbox()
}
