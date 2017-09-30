package gtd_bot

func commandAbort(i *interaction, arguments []string) {
	i.state.setActiveQuestion(questionCollectingInbox)
	askCollectingInbox(i)
}
