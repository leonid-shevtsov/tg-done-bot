package gtd_bot

const questionCollectingInbox = "collecting_inbox"

func init() {
	registerQuestion(questionCollectingInbox, askCollectingInbox, handleCollectingInbox)
}

func askCollectingInbox(i *interaction) {
	if i.state.someWorkToBeDone() {
		i.reply().text(i.locale.CollectingInbox.Prompt).keyboard(
			[][]string{{i.locale.CollectingInbox.StartWorking}},
		).send()
	} else {
		i.sendText(i.locale.CollectingInbox.Prompt)
	}
}

func handleCollectingInbox(i *interaction) string {
	switch i.message.Text {
	case i.locale.CollectingInbox.StartWorking:
		i.state.makeAllContextsActive()
		return nextWorkQuestion(i)
	default:
		i.state.addInboxItem(i.message.Text)
		i.sendText(i.locale.CollectingInbox.Added(i.state.inboxCount()))
		return questionCollectingInbox
	}
}
