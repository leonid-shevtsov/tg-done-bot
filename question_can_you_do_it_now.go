package gtd_bot

const questionCanYouDoItNow = "process_inbox/can_you_do_it_now"

func init() {
	registerQuestion(questionCanYouDoItNow, askCanYouDoItNow, handleCanYouDoItNow)
}

func askCanYouDoItNow(i *interaction) {
	i.sendPrompt(i.locale.CanYouDoItNow.Prompt, [][]string{{
		i.locale.Commands.Yes,
		i.locale.Commands.No,
	}})
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
