package tg_done_bot

import "strings"

func commandWaiting(i *interaction, arguments []string) {
	waitingFors := i.state.allWaitingFors()
	if len(waitingFors) > 0 {
		lines := []string{i.locale.Slash.AllWaitingFors, ""}
		for _, action := range waitingFors {
			lines = append(lines, formatWaitingFor(action), "")
		}
		i.reply().html(strings.Join(lines, "\n")).send()
	} else {
		i.sendText(i.locale.Slash.NoWaitingFors)
	}
	i.askActiveQuestion()
}
