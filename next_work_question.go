package gtd_bot

func nextWorkQuestion(i *interaction) string {
	if i.state.inboxCount() > 0 {
		return questionProcessInbox
	} else if i.state.goalToReviewCount() > 0 {
		return questionReviewGoal
	} else if i.state.goalWithNoActionCount() > 0 {
		i.state.setCurrentGoal(i.state.goalWithNoAction())
		return questionMoveGoalForward
	} else if i.state.waitingForCount() > 0 {
		return questionCheckWaitingFor
	} else if i.state.actionCount() > 0 {
		return questionActionSuggestion
	} else {
		i.sendMessage(i.locale.CollectingInbox.NoMoreWork)
		return questionCollectingInbox
	}
}
