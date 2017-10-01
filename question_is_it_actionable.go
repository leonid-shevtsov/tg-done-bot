package gtd_bot

const questionIsItActionable = "process_inbox/is_it_actionable"
const questionProcessInbox = questionIsItActionable

func init() {
	registerQuestion(questionIsItActionable, askIsItActionable, handleIsItActionable)
}

func askIsItActionable(i *interaction) {
	i.state.startProcessing(i.state.user.CurrentInboxItem)
	i.sendMessage(i.locale.IsItActionable.ProcessingInboxItem)
	i.sendBoldMessage(i.state.user.CurrentInboxItem.Text)
	i.sendPrompt(i.locale.IsItActionable.Prompt, [][]string{{
		i.locale.Commands.Yes,
		i.locale.IsItActionable.NoTrashIt,
		i.locale.Processing.Abort,
	}})
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
