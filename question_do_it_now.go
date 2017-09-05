package gtd_bot

const questionDoItNow = "process_inbox/do_it_now"

func init() {
	registerQuestion(questionDoItNow, askDoItNow, handleDoItNow)
}

func askDoItNow(i *interaction) {
	i.sendPrompt(i.locale.DoItNow.Prompt, [][]string{{
		i.locale.Commands.Done,
		i.locale.Commands.DoItLater,
	}})
	// TODO - notify with timer
	// go i.setDoItNowTimer(i.user.ID, i.user.CurrentActionID)
}

func handleDoItNow(i *interaction) string {
	switch i.message.Text {
	case i.locale.Commands.Done:
		i.state.markActionCompleted()
		return nextWorkQuestion(i)
	case i.locale.Commands.DoItLater:
		return nextWorkQuestion(i)
	default:
		return answerUnclear
	}
}
