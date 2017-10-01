package gtd_bot

import "strings"

func commandInbox(i *interaction, arguments []string) {
	inboxString := strings.Join(arguments, " ")
	i.state.addInboxItem(inboxString)
	i.sendMessage(i.locale.CollectingInbox.Added(i.state.inboxCount()))
	i.askActiveQuestion()
}
