package tg_done_bot

const questionDoItNow = "process_inbox/do_it_now"

func init() {
	registerQuestion(questionDoItNow, askDoItNow, handleDoItNow)
}

func askDoItNow(i *interaction) {
	i.reply().text(i.locale.DoItNow.Prompt).keyboard([][]string{{
		i.locale.Commands.Done,
		i.locale.Commands.DoItLater,
	}}).send()
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
