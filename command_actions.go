package tg_done_bot

import "strings"

func commandActions(i *interaction, arguments []string) {
	actions := i.state.allActions()
	if len(actions) > 0 {
		lines := []string{i.locale.Slash.AllActions, ""}
		for _, action := range actions {
			lines = append(lines, formatAction(action), "")
		}
		i.reply().html(strings.Join(lines, "\n")).send()
	} else {
		i.sendText(i.locale.Slash.NoActions)
	}
	i.askActiveQuestion()
}
