package gtd_bot

import telegram "github.com/go-telegram-bot-api/telegram-bot-api"

const noCommand = "No"

var canYouDoItNowKeyboard = [][]telegram.KeyboardButton{
	[]telegram.KeyboardButton{{Text: yesCommand}, {Text: noCommand}},
}

func (i *interaction) gotoCanYouDoItNow() {
	i.user.State = int(canYouDoItNowState)
	i.repo.update(i.user)
	i.sendPrompt("Can you do it now in 2 minutes?", canYouDoItNowKeyboard)
}

func (i *interaction) handleCanYouDoItNow() {
	switch i.message.Text {
	case yesCommand:
		i.gotoDoItNow()
	case noCommand:
		i.gotoProcessInbox()
	default:
		i.sendUnclear()
		i.gotoCanYouDoItNow()
	}
}
