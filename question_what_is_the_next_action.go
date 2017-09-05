package gtd_bot

const questionWhatIsTheNextAction = "process_inbox/what_is_the_next_action"

func init() {
	registerQuestion(questionWhatIsTheNextAction, askWhatIsTheNextAction, handleWhatIsTheNextAction)
}

func askWhatIsTheNextAction(i *interaction) {
	i.sendPrompt(i.locale.WhatIsTheNextAction.Prompt, [][]string{{
		i.locale.Processing.TrashIt,
		i.locale.Processing.Abort,
	}})
}

func handleWhatIsTheNextAction(i *interaction) string {
	switch i.message.Text {
	case i.locale.Processing.TrashIt:
		return endProcessingByTrashing(i)
	case i.locale.Processing.Abort:
		return endProcessingByAborting(i)
	default:
		i.state.createActionAndMakeCurrent(i.message.Text)
		return questionCanYouDoItNow
	}
}
