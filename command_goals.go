package tg_done_bot

import "strings"

func commandGoals(i *interaction, arguments []string) {
	goals := i.state.allGoals()
	if len(goals) > 0 {
		lines := []string{i.locale.Slash.AllGoals, ""}
		for _, goal := range goals {
			lines = append(lines, formatGoal(i.locale, goal), "")
		}
		i.reply().html(strings.Join(lines, "\n")).send()
	} else {
		i.sendText(i.locale.Slash.NoGoals)
	}
	i.askActiveQuestion()
}
