package gtd_bot

const questionMoveGoalForward = "move_goal_forward"

func init() {
	registerQuestion(questionMoveGoalForward, askMoveGoalForward, handleMoveGoalForward)
}

func askMoveGoalForward(i *interaction) {
	i.sendText(i.locale.MoveGoalForward.Prompt)
	i.reply().goal(i.state.user.CurrentGoal).keyboard([][]string{{
		i.locale.MoveGoalForward.GoalIsAchieved,
		i.locale.Commands.TrashGoal,
		i.locale.Commands.WaitingFor,
	}}).send()
}

func handleMoveGoalForward(i *interaction) string {
	switch i.message.Text {
	case i.locale.MoveGoalForward.GoalIsAchieved:
		i.sendText(i.locale.MoveGoalForward.CongratulationsComplete)
		i.state.completeCurrentGoal()
		return nextWorkQuestion(i)
	case i.locale.Commands.TrashGoal:
		i.state.dropCurrentGoal()
		i.sendText(i.locale.Messages.GoalTrashed)
		return nextWorkQuestion(i)
	case i.locale.Commands.WaitingFor:
		return questionWhatIsTheGoalWaitingFor
	default:
		i.state.createActionAndMakeCurrent(i.message.Text)
		i.sendText(i.locale.MoveGoalForward.AddedAction)
		return questionProcessingWhatIsTheContext
	}
}
