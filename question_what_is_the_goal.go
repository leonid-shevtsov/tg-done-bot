package gtd_bot

const questionWhatIsTheGoal = "process_inbox/what_is_the_goal"

func init() {
	registerQuestion(questionWhatIsTheGoal, askWhatIsTheGoal, handleWhatIsTheGoal)
}

func askWhatIsTheGoal(i *interaction) {
	i.reply().text(i.locale.WhatIsTheGoal.Prompt).keyboard([][]string{{
		i.locale.Processing.TrashIt,
		i.locale.Processing.Abort,
	}}).send()
}

func handleWhatIsTheGoal(i *interaction) string {
	switch i.message.Text {
	case i.locale.Processing.TrashIt:
		return endProcessingByTrashing(i)
	case i.locale.Processing.Abort:
		return endProcessingByAborting(i)
	default:
		i.state.createGoalAndMakeCurrent(i.message.Text)
		return questionWhatIsTheDueDate
	}
}
