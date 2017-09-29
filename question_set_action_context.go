package gtd_bot

const questionSetActionContext = "set_action_context"

func init() {
	registerQuestion(questionSetActionContext, askSetActionContext, handleSetActionContext)
}

func askSetActionContext(i *interaction) {
	fullKeyboard := append([][]string{
		{
			i.locale.WhatIsTheContext.None,
			i.locale.WhatIsTheContext.NewContext,
			i.locale.Commands.Keep,
		},
	}, buildContextKeyboard(i)...)
	i.sendPrompt(i.locale.WhatIsTheContext.Prompt, fullKeyboard)
}

func handleSetActionContext(i *interaction) string {
	switch i.message.Text {
	case i.locale.WhatIsTheContext.None:
		i.state.setCurrentActionContext(nil)
		i.sendMessage(i.locale.WhatIsTheContext.Success)
		return nextWorkQuestion(i)
	case i.locale.WhatIsTheContext.NewContext:
		return questionCreateContext
	case i.locale.Commands.Keep:
		return nextWorkQuestion(i)
	default:
		context := i.state.findContextByText(i.message.Text)
		if context != nil {
			i.state.setCurrentActionContext(context)
			i.state.markCurrentContextInactive()
			i.sendMessage(i.locale.WhatIsTheContext.Success)
			i.sendMessage(i.locale.WhatIsTheContext.MarkingInactive)
			return nextWorkQuestion(i)
		} else {
			return answerUnclear
		}
	}
}
