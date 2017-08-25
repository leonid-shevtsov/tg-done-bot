package gtd_bot

import (
	"fmt"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

const trashGoalCommand = "Let's just trash it"

var whatIsTheGoalKeyboard = [][]telegram.KeyboardButton{
	[]telegram.KeyboardButton{{Text: trashGoalCommand}, {Text: abortCommand}},
}

func (i *interaction) gotoWhatIsTheGoal() {
	i.user.State = int(whatIsTheGoalState)
	i.repo.update(i.user)
	i.sendPrompt("What is the goal here?", whatIsTheGoalKeyboard)
}

func (i *interaction) handleWhatIsTheGoal() {
	switch i.message.Text {
	case trashGoalCommand:
		i.trashCurrentInboxItem()
	case abortCommand:
		i.abortProcessing()
	default:
		i.createGoal()
		i.gotoWhatIsTheNextAction()
	}
}

func (i *interaction) createGoal() {
	goal := &Goal{
		UserID: i.user.ID,
		Text:   i.message.Text,
	}
	i.repo.insert(goal)
	i.user.CurrentGoalID = goal.ID
	i.repo.update(i.user)
}

func (i *interaction) abortProcessing() {
	abortText := fmt.Sprintf("OK. Items left to process: %d.", i.inboxCount())
	i.sendMessage(abortText)
	i.gotoInitialState()
}
