package tg_done_bot

const questionProcessingCreateContext = "process_inbox/create_context"

func init() {
	registerQuestion(questionProcessingCreateContext, askProcessingCreateContext, handleProcessingCreateContext)
}

func askProcessingCreateContext(i *interaction) {
	i.reply().text(i.locale.CreateContext.Prompt).keyboard([][]string{
		{i.locale.CreateContext.Cancel},
	}).send()
}

func handleProcessingCreateContext(i *interaction) string {
	switch i.message.Text {
	case i.locale.CreateContext.Cancel:
		return questionProcessingWhatIsTheContext
	default:
		context, err := i.state.createContext(i.message.Text)
		if err != nil {
			i.sendText(i.locale.CreateContext.AlreadyExists)
			return questionProcessingCreateContext
		}
		i.state.setCurrentActionContext(context)
		i.sendText(i.locale.CreateContext.Success)
		return nextWorkQuestion(i)
	}
}
