package gtd_bot

func nextWorkQuestion(i *interaction) string {
	if inboxItemToProcess := i.state.inboxItemToProcess(); inboxItemToProcess != nil {
		i.state.startProcessing(inboxItemToProcess)
		return questionProcessInbox
	} else if goal := i.state.goalToReview(); goal != nil {
		i.state.setCurrentGoal(goal)
		return questionReviewGoal
	} else if i.state.goalWithNoActionCount() > 0 {
		i.state.setCurrentGoal(i.state.goalWithNoAction())
		return questionMoveGoalForward
	} else if waitingFor := i.state.waitingForToCheck(); waitingFor != nil {
		i.state.setCurrentWaitingFor(waitingFor)
		return questionCheckWaitingFor
	} else if actionToDo := i.state.actionToDo(); actionToDo != nil {
		i.state.setSuggestedAction(actionToDo)
		return questionActionSuggestion
	} else {
		i.sendText(i.locale.CollectingInbox.NoMoreWork)
		return questionCollectingInbox
	}
}
