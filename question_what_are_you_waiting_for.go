package gtd_bot

const questionWhatAreYouWaitingFor = "process_inbox/waiting_for"

func init() {
	registerQuestion(questionWhatAreYouWaitingFor, askWhatAreYouWaitingFor, handleWhatAreYouWaitingFor)
}

func askWhatAreYouWaitingFor(i *interaction) {
	i.reply().text(i.locale.WhatAreYouWaitingFor.Prompt).keyboard([][]string{
		{i.locale.Processing.TrashIt},
		{i.locale.WhatAreYouWaitingFor.Nothing},
		{i.locale.Processing.Abort},
	}).send()
}

func handleWhatAreYouWaitingFor(i *interaction) string {
	switch i.message.Text {
	case i.locale.Processing.TrashIt:
		return endProcessingByTrashing(i)
	case i.locale.WhatAreYouWaitingFor.Nothing:
		return questionWhatIsTheNextAction
	case i.locale.Processing.Abort:
		return endProcessingByAborting(i)
	default:
		i.state.createWaitingForAndMakeCurrent(i.message.Text)
		return nextWorkQuestion(i)
	}
}
