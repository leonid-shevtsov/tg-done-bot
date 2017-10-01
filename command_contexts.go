package gtd_bot

import (
	"fmt"
	"strings"
)

func commandContexts(i *interaction, arguments []string) {
	contexts := i.state.allContexts()
	if len(contexts) > 0 {
		lines := []string{i.locale.Slash.AllContexts, ""}
		for _, context := range contexts {
			contextText := formatContext(context)
			actionCount := i.state.actionInContextCount(context)
			line := fmt.Sprintf("%s (%s)", contextText, i.locale.Plurals.Actions(actionCount))
			lines = append(lines, line, "")
		}
		i.reply().html(strings.Join(lines, "\n")).send()
	} else {
		i.sendText(i.locale.Slash.NoContexts)
	}
	i.askActiveQuestion()
}
