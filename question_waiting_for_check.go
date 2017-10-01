package gtd_bot

const questionCheckWaitingFor = "check_waiting_for"

func init() {
	registerQuestion(questionCheckWaitingFor, askCheckWaitingFor, handleCheckWaitingFor)
}

func askCheckWaitingFor(i *interaction) {
	i.sendText(i.locale.CheckWaitingFor.YourGoal)
	i.reply().goal(i.state.user.CurrentWaitingFor.Goal).send()
	i.sendText(i.locale.CheckWaitingFor.IsWaitingFor)
	i.reply().waitingFor(i.state.user.CurrentWaitingFor).keyboard([][]string{
		{i.locale.CheckWaitingFor.ItIsReady},
		{i.locale.CheckWaitingFor.StillWaiting},
		{i.locale.Commands.TrashGoal},
		{i.locale.Commands.BackToInbox},
	}).send()
}

func handleCheckWaitingFor(i *interaction) string {
	switch i.message.Text {
	case i.locale.CheckWaitingFor.ItIsReady:
		i.state.completeCurrentWaitingFor()
		i.sendText(i.locale.CheckWaitingFor.Success)
		return questionMoveGoalForward
	case i.locale.CheckWaitingFor.StillWaiting:
		i.state.continueToWait()
		i.sendText(i.locale.CheckWaitingFor.ContinuingToWait)
		return nextWorkQuestion(i)
	case i.locale.Commands.TrashGoal:
		i.state.dropCurrentGoal()
		i.sendText(i.locale.Messages.GoalTrashed)
		return nextWorkQuestion(i)
	case i.locale.Commands.BackToInbox:
		return questionCollectingInbox
	default:
		return answerUnclear
	}
}
