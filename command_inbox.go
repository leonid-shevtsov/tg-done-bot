package gtd_bot

import "strings"

func commandInbox(i *interaction, arguments []string) {
	if len(arguments) > 0 {
		inboxString := strings.Join(arguments, " ")
		i.state.addInboxItem(inboxString)
		i.reply().text(i.locale.CollectingInbox.Added(i.state.inboxCount())).send()
	} else {
		inboxItems := i.state.allInboxItems()
		if len(inboxItems) > 0 {
			lines := []string{i.locale.Slash.AllInboxItems, ""}
			for _, inboxItem := range inboxItems {
				lines = append(lines, formatInboxItem(inboxItem), "")
			}
			i.reply().html(strings.Join(lines, "\n")).send()
		} else {
			i.sendText(i.locale.Slash.NoInboxItems)
		}
	}
	i.askActiveQuestion()
}
