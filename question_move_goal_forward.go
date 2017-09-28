package gtd_bot

const questionMoveGoalForward = "move_goal_forward"

func init() {
	registerQuestion(questionMoveGoalForward, askMoveGoalForward, handleMoveGoalForward)
}

func askMoveGoalForward(i *interaction) {
	i.sendMessage(i.locale.MoveGoalForward.Prompt)
	i.sendBoldPrompt(i.state.user.CurrentGoal.Text, [][]string{{
		i.locale.MoveGoalForward.GoalIsAchieved,
		i.locale.Commands.TrashGoal,
		i.locale.Commands.WaitingFor,
	}})
}

func handleMoveGoalForward(i *interaction) string {
	switch i.message.Text {
	case i.locale.MoveGoalForward.GoalIsAchieved:
		i.sendMessage(i.locale.MoveGoalForward.CongratulationsComplete)
		i.state.completeCurrentGoal()
		return nextWorkQuestion(i)
	case i.locale.Commands.TrashGoal:
		i.state.dropCurrentGoal()
		i.sendMessage(i.locale.Messages.GoalTrashed)
		return nextWorkQuestion(i)
	case i.locale.Commands.WaitingFor:
		return questionWhatIsTheGoalWaitingFor
	default:
		i.state.createActionAndMakeCurrent(i.message.Text)
		i.sendMessage(i.locale.MoveGoalForward.AddedAction)
		return nextWorkQuestion(i)
	}
}
