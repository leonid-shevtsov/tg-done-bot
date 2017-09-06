package gtd_bot

const questionActionSuggestion = "action_suggestion"

func init() {
	registerQuestion(questionActionSuggestion, askActionSuggestion, handleActionSuggestion)
}

func askActionSuggestion(i *interaction) {
	if actionToDo := i.state.actionToDo(); actionToDo != nil {
		i.state.setSuggestedAction(actionToDo)
		i.sendMessage(i.locale.ActionSuggestion.IThinkYouShouldWorkOn)
		i.sendMessage(actionToDo.Goal.Text)
		i.sendMessage(i.locale.ActionSuggestion.ByDoing)
		i.sendPrompt(actionToDo.Text, [][]string{
			{i.locale.ActionSuggestion.Doing},
			{i.locale.ActionSuggestion.Skip},
			{i.locale.ActionSuggestion.ItIsDone},
			{i.locale.ActionSuggestion.ChangeNextAction},
			{i.locale.Commands.WaitingFor},
			{i.locale.Commands.TrashGoal},
			{i.locale.Commands.BackToInbox},
		})
	} else {
		panic("bad precondition for action_suggestion question")
	}
}

func handleActionSuggestion(i *interaction) string {
	switch i.message.Text {
	case i.locale.ActionSuggestion.Doing:
		return questionDoing
	case i.locale.ActionSuggestion.Skip:
		i.state.skipCurrentAction()
		i.sendMessage(i.locale.ActionSuggestion.Skipping)
		return nextWorkQuestion(i)
	case i.locale.ActionSuggestion.ItIsDone:
		i.state.completeCurrentAction()
		i.sendMessage(i.locale.Doing.Completed)
		return questionMoveGoalForward
	case i.locale.ActionSuggestion.ChangeNextAction:
		i.state.dropCurrentAction()
		return questionMoveGoalForward
	case i.locale.Commands.WaitingFor:
		i.state.dropCurrentAction()
		return questionWhatIsTheGoalWaitingFor
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
