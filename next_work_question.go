package gtd_bot

func nextWorkQuestion(i *interaction) string {
	if i.state.inboxCount() > 0 {
		return questionProcessInbox
	} else if i.state.waitingForCount() > 0 {
		return questionCheckWaitingFor
	} else if i.state.actionCount() > 0 {
		return questionActionSuggestion
	} else {
		i.sendMessage(i.locale.CollectingInbox.NoMoreWork)
		return questionCollectingInbox
	}
}
