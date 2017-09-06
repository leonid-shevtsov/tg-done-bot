package gtd_bot

const questionCollectingInbox = "collecting_inbox"

func init() {
	registerQuestion(questionCollectingInbox, askCollectingInbox, handleCollectingInbox)
}

func askCollectingInbox(i *interaction) {
	if i.state.someWorkToBeDone() {
		i.sendPrompt(i.locale.CollectingInbox.Prompt, [][]string{{i.locale.CollectingInbox.StartWorking}})
	} else {
		i.sendMessage(i.locale.CollectingInbox.Prompt)
	}
}

func handleCollectingInbox(i *interaction) string {
	switch i.message.Text {
	case i.locale.CollectingInbox.StartWorking:
		return nextWorkQuestion(i)
	default:
		i.state.addInboxItem(i.message.Text)
		i.sendMessage(i.locale.CollectingInbox.Added(i.state.inboxCount()))
		return questionCollectingInbox
	}
}
