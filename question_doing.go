package gtd_bot

const questionDoing = "doing"

func init() {
	registerQuestion(questionDoing, askDoing, handleDoing)
}

func askDoing(i *interaction) {
	i.reply().text(i.locale.Doing.Prompt).keyboard([][]string{{
		i.locale.Commands.Done,
		i.locale.Commands.DoItLater,
	}}).send()
}

func handleDoing(i *interaction) string {
	switch i.message.Text {
	case i.locale.Commands.Done:
		i.state.completeCurrentAction()
		i.sendText(i.locale.Doing.Completed)
		return questionMoveGoalForward
	case i.locale.Commands.DoItLater:
		i.state.skipCurrentAction()
		return nextWorkQuestion(i)
	default:
		return answerUnclear
	}
}
