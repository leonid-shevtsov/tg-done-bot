package tg_done_bot

const questionActionSuggestion = "action_suggestion"

func init() {
	registerQuestion(questionActionSuggestion, askActionSuggestion, handleActionSuggestion)
}

func askActionSuggestion(i *interaction) {
	i.sendText(i.locale.ActionSuggestion.IThinkYouShouldWorkOn)
	i.reply().goal(i.state.user.CurrentAction.Goal).send()
	i.sendText(i.locale.ActionSuggestion.ByDoing)

	var contextAction string
	if i.state.user.CurrentAction.ContextID != 0 {
		contextAction = i.locale.ActionSuggestion.WrongContext
	} else {
		contextAction = i.locale.ActionSuggestion.NeedContext
	}

	i.reply().action(i.state.user.CurrentAction).keyboard([][]string{
		{
			i.locale.ActionSuggestion.Doing,
			i.locale.ActionSuggestion.Skip,
			i.locale.ActionSuggestion.ItIsDone,
		},
		{
			i.locale.ActionSuggestion.ChangeNextAction,
			i.locale.Commands.WaitingFor,
			contextAction,
		},
		{
			i.locale.Commands.TrashGoal,
			i.locale.Commands.BackToInbox,
		},
	}).send()
}

func handleActionSuggestion(i *interaction) string {
	switch i.message.Text {
	case i.locale.ActionSuggestion.Doing:
		return questionDoing
	case i.locale.ActionSuggestion.Skip:
		i.state.skipCurrentAction()
		i.sendText(i.locale.ActionSuggestion.Skipping)
		return nextWorkQuestion(i)
	case i.locale.ActionSuggestion.ItIsDone:
		i.state.completeCurrentAction()
		i.sendText(i.locale.Doing.Completed)
		return questionMoveGoalForward
	case i.locale.ActionSuggestion.ChangeNextAction:
		i.state.dropCurrentAction()
		return questionMoveGoalForward
	case i.locale.Commands.WaitingFor:
		i.state.dropCurrentAction()
		return questionWhatIsTheGoalWaitingFor
	case i.locale.Commands.TrashGoal:
		i.state.dropCurrentGoal()
		i.sendText(i.locale.Messages.GoalTrashed)
		return nextWorkQuestion(i)
	case i.locale.Commands.BackToInbox:
		return questionCollectingInbox
	case i.locale.ActionSuggestion.WrongContext:
		i.state.markCurrentContextInactive()
		i.sendText(i.locale.ActionSuggestion.ContextNowInactive)
		return nextWorkQuestion(i)
	case i.locale.ActionSuggestion.NeedContext:
		return questionSetActionContext
	default:
		return answerUnclear
	}
}
