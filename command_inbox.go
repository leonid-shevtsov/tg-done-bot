package gtd_bot

import "strings"

func commandInbox(i *interaction, arguments []string) {
	inboxString := strings.Join(arguments, " ")
	i.state.addInboxItem(inboxString)
	i.reply().text(i.locale.CollectingInbox.Added(i.state.inboxCount())).send()
	i.askActiveQuestion()
}
