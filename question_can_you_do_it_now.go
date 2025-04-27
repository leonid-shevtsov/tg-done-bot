package tg_done_bot

const questionCanYouDoItNow = "process_inbox/can_you_do_it_now"

func init() {
	registerQuestion(questionCanYouDoItNow, askCanYouDoItNow, handleCanYouDoItNow)
}

func askCanYouDoItNow(i *interaction) {
	i.reply().text(i.locale.CanYouDoItNow.Prompt).keyboard([][]string{{
		i.locale.Commands.Yes,
		i.locale.Commands.No,
	}}).send()
}

func handleCanYouDoItNow(i *interaction) string {
	switch i.message.Text {
	case i.locale.Commands.Yes:
		return questionDoItNow
	case i.locale.Commands.No:
		return questionProcessingWhatIsTheContext
	default:
		return answerUnclear
	}
}
