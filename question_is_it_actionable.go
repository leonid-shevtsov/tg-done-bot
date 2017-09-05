package gtd_bot

const questionIsItActionable = "process_inbox/is_it_actionable"
const questionProcessInbox = questionIsItActionable

func init() {
	registerQuestion(questionIsItActionable, askIsItActionable, handleIsItActionable)
}

func askIsItActionable(i *interaction) {
	if inboxItemToProcess := i.state.inboxItemToProcess(); inboxItemToProcess != nil {
		i.state.startProcessing(inboxItemToProcess)
		i.sendMessage(i.locale.IsItActionable.ProcessingInboxItem)
		i.sendMessage(inboxItemToProcess.Text)
		i.sendPrompt(i.locale.IsItActionable.Prompt, [][]string{{
			i.locale.Commands.Yes,
			i.locale.IsItActionable.NoTrashIt,
			i.locale.Processing.Abort,
		}})
	} else {
		panic("bad question prerequisites for is it actionable")
	}
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
