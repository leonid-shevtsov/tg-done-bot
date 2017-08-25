package gtd_bot

import (
	"fmt"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

const processInboxCommand = "Process inbox"

var initialStateKeyboard = [][]telegram.KeyboardButton{
	[]telegram.KeyboardButton{{Text: processInboxCommand}},
}

func (i *interaction) gotoInitialState() {
	i.user.State = int(initialState)
	i.user.CurrentInboxItemID = 0
	i.user.CurrentGoalID = 0
	i.repo.update(i.user)

	if i.inboxCount() > 0 {
		i.sendPrompt("Back to collecting inbox.", initialStateKeyboard)
	} else {
		i.sendMessage("Back to collecting inbox.")
	}
}

func (i *interaction) handleInitial() {
	switch i.message.Text {
	case processInboxCommand:
		i.gotoProcessInbox()
	default:
		i.addInboxItem()
	}
}

func (i *interaction) addInboxItem() {
	inboxItem := &InboxItem{
		UserID: i.user.ID,
		Text:   i.message.Text,
	}
	i.repo.insert(inboxItem)
	responseText := fmt.Sprintf("Added to inbox. Now in inbox: %d items.", i.inboxCount())
	i.sendPrompt(responseText, initialStateKeyboard)
}
