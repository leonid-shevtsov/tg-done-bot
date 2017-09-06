package gtd_bot

const questionCheckWaitingFor = "check_waiting_for"

func init() {
	registerQuestion(questionCheckWaitingFor, askCheckWaitingFor, handleCheckWaitingFor)
}

func askCheckWaitingFor(i *interaction) {
	if waitingFor := i.state.waitingForToCheck(); waitingFor != nil {
		i.state.setCurrentWaitingFor(waitingFor)
		i.sendMessage(i.locale.CheckWaitingFor.YourGoal)
		i.sendMessage(waitingFor.Goal.Text)
		i.sendMessage(i.locale.CheckWaitingFor.IsWaitingFor)
		i.sendPrompt(waitingFor.Text, [][]string{
			{i.locale.CheckWaitingFor.ItIsReady},
			{i.locale.CheckWaitingFor.StillWaiting},
			{i.locale.Commands.TrashGoal},
			{i.locale.Commands.BackToInbox},
		})
	} else {
		panic("bad precondition for waiting_for_check question")
	}
}

func handleCheckWaitingFor(i *interaction) string {
	switch i.message.Text {
	case i.locale.CheckWaitingFor.ItIsReady:
		i.state.completeCurrentWaitingFor()
		i.sendMessage(i.locale.CheckWaitingFor.Success)
		return questionMoveGoalForward
	case i.locale.CheckWaitingFor.StillWaiting:
		i.state.continueToWait()
		i.sendMessage(i.locale.CheckWaitingFor.ContinuingToWait)
		return nextWorkQuestion(i)
	case i.locale.Commands.TrashGoal:
		i.state.dropCurrentGoal()
		i.sendMessage(i.locale.Messages.GoalTrashed)
		return nextWorkQuestion(i)
	case i.locale.Commands.BackToInbox:
		return questionCollectingInbox
	default:
		return answerUnclear
	}
}
