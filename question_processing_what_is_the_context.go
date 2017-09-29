package gtd_bot

const questionProcessingWhatIsTheContext = "process_inbox/what_is_the_context"

func init() {
	registerQuestion(questionProcessingWhatIsTheContext, askProcessingWhatIsTheContext, handleProcessingWhatIsTheContext)
}

func askProcessingWhatIsTheContext(i *interaction) {
	fullKeyboard := append([][]string{
		{
			i.locale.WhatIsTheContext.None,
			i.locale.WhatIsTheContext.NewContext,
		},
	}, buildContextKeyboard(i)...)
	i.sendPrompt(i.locale.WhatIsTheContext.Prompt, fullKeyboard)
}

func handleProcessingWhatIsTheContext(i *interaction) string {
	switch i.message.Text {
	case i.locale.WhatIsTheContext.None:
		i.state.setCurrentActionContext(nil)
		i.sendMessage(i.locale.WhatIsTheContext.Success)
		return nextWorkQuestion(i)
	case i.locale.WhatIsTheContext.NewContext:
		return questionProcessingCreateContext
	default:
		context := i.state.findContextByText(i.message.Text)
		if context != nil {
			i.state.setCurrentActionContext(context)
			i.sendMessage(i.locale.WhatIsTheContext.Success)
			return nextWorkQuestion(i)
		} else {
			return answerUnclear
		}
	}
}
