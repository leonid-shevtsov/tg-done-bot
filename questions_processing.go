package gtd_bot

func endProcessingByTrashing(i *interaction) string {
	i.state.trashCurrentInboxItem()
	i.sendText(i.locale.IsItActionable.Trashed)
	return nextWorkQuestion(i)
}

func endProcessingByAborting(i *interaction) string {
	i.sendText(i.locale.Processing.Aborted(i.state.inboxCount()))
	return questionCollectingInbox
}
