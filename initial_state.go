package gtd_bot

import (
	"fmt"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

const beProductiveCommand = "I'm ready for some work"

var initialStateKeyboard = [][]telegram.KeyboardButton{
	[]telegram.KeyboardButton{{Text: beProductiveCommand}},
}

func (i *interaction) gotoInitialState() {
	previousState := interactionState(i.user.State)
	i.user.State = int(initialState)
	i.user.CurrentInboxItemID = 0
	i.user.CurrentGoalID = 0
	i.repo.update(i.user)

	if previousState == onboardingState {
		i.sendMessage("Collecting inbox now. Send me anything that comes up, one thing at a time.")
	} else if i.someWorkToBeDone() {
		i.sendPrompt("Back to collecting inbox.", initialStateKeyboard)
	} else {
		i.sendMessage("Back to collecting inbox.")
	}
}

func (i *interaction) someWorkToBeDone() bool {
	return i.inboxCount() > 0 || i.actionCount() > 0
}

func (i *interaction) handleInitial() {
	switch i.message.Text {
	case beProductiveCommand:
		i.gotoNextWorkUnit()
	default:
		i.addInboxItem()
	}
}

func (i *interaction) gotoNextWorkUnit() {
	if i.inboxCount() > 0 {
		i.gotoProcessInbox()
	} else if i.actionCount() > 0 {
		i.gotoActionSuggestionState()
	} else {
		i.sendMessage("No more work!")
		i.gotoInitialState()
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
