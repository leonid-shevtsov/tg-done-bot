package tg_done_bot

func buildContextKeyboard(i *interaction) [][]string {
	contexts := i.state.allContexts()
	var contextKeyboard [][]string
	for _, context := range contexts {
		lastRowIndex := len(contextKeyboard) - 1
		contextLabel := context.Text
		if lastRowIndex >= 0 && len(contextKeyboard[lastRowIndex]) < 3 {
			contextKeyboard[lastRowIndex] = append(contextKeyboard[lastRowIndex], contextLabel)
		} else {
			contextKeyboard = append(contextKeyboard, []string{contextLabel})
		}
	}
	return contextKeyboard
}
