package tg_done_bot

const questionWhatIsTheNextAction = "process_inbox/what_is_the_next_action"

func init() {
	registerQuestion(questionWhatIsTheNextAction, askWhatIsTheNextAction, handleWhatIsTheNextAction)
}

func askWhatIsTheNextAction(i *interaction) {
	i.reply().text(i.locale.WhatIsTheNextAction.Prompt).keyboard([][]string{
		{i.locale.Processing.TrashIt},
		{i.locale.WhatIsTheNextAction.WaitingFor},
		{i.locale.Processing.Abort},
	}).send()
}

func handleWhatIsTheNextAction(i *interaction) string {
	switch i.message.Text {
	case i.locale.Processing.TrashIt:
		return endProcessingByTrashing(i)
	case i.locale.WhatIsTheNextAction.WaitingFor:
		return questionWhatAreYouWaitingFor
	case i.locale.Processing.Abort:
		return endProcessingByAborting(i)
	default:
		i.state.createActionAndMakeCurrent(i.message.Text)
		return questionCanYouDoItNow
	}
}
