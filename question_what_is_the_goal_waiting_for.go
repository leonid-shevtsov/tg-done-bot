package gtd_bot

const questionWhatIsTheGoalWaitingFor = "doing/waiting_for"

func init() {
	registerQuestion(questionWhatIsTheGoalWaitingFor, askWhatIsTheGoalWaitingFor, handleWhatIsTheGoalWaitingFor)
}

func askWhatIsTheGoalWaitingFor(i *interaction) {
	i.sendPrompt(i.locale.WhatIsTheGoalWaitingFor.Prompt, [][]string{
		{i.locale.WhatIsTheGoalWaitingFor.Nothing},
	})
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
