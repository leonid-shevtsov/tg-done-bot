package gtd_bot

const questionProcessingCreateContext = "process_inbox/create_context"

func init() {
	registerQuestion(questionProcessingCreateContext, askProcessingCreateContext, handleProcessingCreateContext)
}

func askProcessingCreateContext(i *interaction) {
	i.sendPrompt(i.locale.CreateContext.Prompt, [][]string{
		{i.locale.CreateContext.Cancel},
	})
}

func handleProcessingCreateContext(i *interaction) string {
	switch i.message.Text {
	case i.locale.CreateContext.Cancel:
		return questionProcessingWhatIsTheContext
	default:
		context, err := i.state.createContext(i.message.Text)
		if err != nil {
			i.sendMessage(i.locale.CreateContext.AlreadyExists)
			return questionProcessingCreateContext
		}
		i.state.setCurrentActionContext(context)
		i.sendMessage(i.locale.CreateContext.Success)
		return nextWorkQuestion(i)
	}
}
