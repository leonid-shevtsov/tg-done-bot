package gtd_bot

import (
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

var whatIsTheNextActionKeyboard = [][]telegram.KeyboardButton{
	[]telegram.KeyboardButton{{Text: trashGoalCommand}, {Text: abortCommand}},
}

func (i *interaction) gotoWhatIsTheNextAction() {
	i.user.State = int(whatIsTheNextActionState)
	i.repo.update(i.user)
	i.sendPrompt("What is the next physical action?", whatIsTheNextActionKeyboard)
}

func (i *interaction) handleWhatIsTheNextAction() {
	switch i.message.Text {
	case trashGoalCommand:
		i.trashCurrentInboxItem()
	case abortCommand:
		i.abortProcessing()
	default:
		i.createAction()
		i.gotoCanYouDoItNow()
	}
}

func (i *interaction) createAction() {
	action := &Action{
		UserID: i.user.ID,
		GoalID: i.user.CurrentGoalID,
		Text:   i.message.Text,
	}
	i.repo.insert(action)
	i.user.CurrentActionID = action.ID
	i.repo.update(i.user)
	// the inbox item is considered processed once there is an action
	i.user.CurrentInboxItem.ProcessedAt = time.Now()
	i.repo.update(i.user.CurrentInboxItem)
}
