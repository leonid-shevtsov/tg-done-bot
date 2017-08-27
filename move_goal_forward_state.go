package gtd_bot

import (
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

const goalIsAchievedCommand = "This goal is achieved"
const reviewItLaterCommand = "Let's review it later"

var moveGoalForwardKeyboard = [][]telegram.KeyboardButton{
	[]telegram.KeyboardButton{{Text: goalIsAchievedCommand}, {Text: reviewItLaterCommand}},
}

func (i *interaction) gotoMoveGoalForward() {
	i.user.State = int(moveGoalForwardState)
	i.repo.update(i.user)

	i.sendMessage("Now you are closer to achieving:")
	i.sendMessage(i.user.CurrentGoal.Text)
	i.sendPrompt("What is the next action towards this goal?", moveGoalForwardKeyboard)
}

func (i *interaction) handleMoveGoalForward() {
	switch i.message.Text {
	case goalIsAchievedCommand:
		i.sendMessage("Congratulations on succeeding!")
		i.markGoalAsCompleted()
		i.gotoNextWorkUnit()
	case reviewItLaterCommand:
		i.sendMessage("OK, marked goal for review.")
		i.gotoNextWorkUnit()
	default:
		i.createAction()
		i.sendMessage("OK, recorded next action for this goal.")
		i.gotoNextWorkUnit()
	}
}

func (i *interaction) markGoalAsCompleted() {
	i.user.CurrentGoal.CompletedAt = time.Now()
	i.repo.update(i.user.CurrentGoal)
}
