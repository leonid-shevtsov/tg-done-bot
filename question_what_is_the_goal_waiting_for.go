package tg_done_bot

const questionWhatIsTheGoalWaitingFor = "doing/waiting_for"

func init() {
	registerQuestion(questionWhatIsTheGoalWaitingFor, askWhatIsTheGoalWaitingFor, handleWhatIsTheGoalWaitingFor)
}

func askWhatIsTheGoalWaitingFor(i *interaction) {
	i.reply().text(i.locale.WhatIsTheGoalWaitingFor.Prompt).keyboard([][]string{
		{i.locale.WhatIsTheGoalWaitingFor.Nothing},
	}).send()
}

func handleWhatIsTheGoalWaitingFor(i *interaction) string {
	switch i.message.Text {
	case i.locale.WhatIsTheGoalWaitingFor.Nothing:
		return questionMoveGoalForward
	default:
		i.state.createWaitingForAndMakeCurrent(i.message.Text)
		return nextWorkQuestion(i)
	}
}
