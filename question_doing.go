package gtd_bot

const questionDoing = "doing"

func init() {
	registerQuestion(questionDoing, askDoing, handleDoing)
}

func askDoing(i *interaction) {
	i.sendPrompt(i.locale.Doing.Prompt, [][]string{{
		i.locale.Commands.Done,
		i.locale.Commands.DoItLater,
	}})
}

func handleDoing(i *interaction) string {
	switch i.message.Text {
	case i.locale.Commands.Done:
		i.state.completeCurrentAction()
		i.sendMessage(i.locale.Doing.Completed)
		return questionMoveGoalForward
	case i.locale.Commands.DoItLater:
		i.state.skipCurrentAction()
		return nextWorkQuestion(i)
	default:
		return answerUnclear
	}
}
