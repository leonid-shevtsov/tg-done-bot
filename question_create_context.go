package gtd_bot

const questionCreateContext = "create_context"

func init() {
	registerQuestion(questionCreateContext, askCreateContext, handleCreateContext)
}

func askCreateContext(i *interaction) {
	i.reply().text(i.locale.CreateContext.Prompt).keyboard([][]string{
		{i.locale.CreateContext.Cancel},
	}).send()
}

func handleCreateContext(i *interaction) string {
	switch i.message.Text {
	case i.locale.CreateContext.Cancel:
		return questionSetActionContext
	default:
		context, err := i.state.createContext(i.message.Text)
		if err != nil {
			i.sendText(i.locale.CreateContext.AlreadyExists)
			return questionCreateContext
		}
		i.state.setCurrentActionContext(context)
		i.sendText(i.locale.CreateContext.Success)
		i.state.markCurrentContextInactive()
		i.sendText(i.locale.WhatIsTheContext.MarkingInactive)
		return nextWorkQuestion(i)
	}
}
