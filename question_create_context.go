package gtd_bot

const questionCreateContext = "create_context"

func init() {
	registerQuestion(questionCreateContext, askCreateContext, handleCreateContext)
}

func askCreateContext(i *interaction) {
	i.sendPrompt(i.locale.CreateContext.Prompt, [][]string{
		{i.locale.CreateContext.Cancel},
	})
}

func handleCreateContext(i *interaction) string {
	switch i.message.Text {
	case i.locale.CreateContext.Cancel:
		return questionSetActionContext
	default:
		context, err := i.state.createContext(i.message.Text)
		if err != nil {
			i.sendMessage(i.locale.CreateContext.AlreadyExists)
			return questionCreateContext
		}
		i.state.setCurrentActionContext(context)
		i.sendMessage(i.locale.CreateContext.Success)
		i.state.markCurrentContextInactive()
		i.sendMessage(i.locale.WhatIsTheContext.MarkingInactive)
		return nextWorkQuestion(i)
	}
}
