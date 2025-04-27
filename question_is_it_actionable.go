package tg_done_bot

const questionIsItActionable = "process_inbox/is_it_actionable"
const questionProcessInbox = questionIsItActionable

func init() {
	registerQuestion(questionIsItActionable, askIsItActionable, handleIsItActionable)
}

func askIsItActionable(i *interaction) {
	i.state.startProcessing(i.state.user.CurrentInboxItem)
	i.sendText(i.locale.IsItActionable.ProcessingInboxItem)
	i.reply().inboxItem(i.state.user.CurrentInboxItem).send()
	i.reply().text(i.locale.IsItActionable.Prompt).keyboard([][]string{{
		i.locale.Commands.Yes,
		i.locale.IsItActionable.NoTrashIt,
		i.locale.Processing.Abort,
	}}).send()
}

func handleIsItActionable(i *interaction) string {
	switch i.message.Text {
	case i.locale.Commands.Yes:
		return questionWhatIsTheGoal
	case i.locale.IsItActionable.NoTrashIt:
		return endProcessingByTrashing(i)
	case i.locale.Processing.Abort:
		return endProcessingByAborting(i)
	default:
		return answerUnclear
	}
}
