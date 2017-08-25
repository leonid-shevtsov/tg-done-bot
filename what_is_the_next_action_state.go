package gtd_bot

import telegram "github.com/go-telegram-bot-api/telegram-bot-api"

var whatIsTheNextActionKeyboard = [][]telegram.KeyboardButton{
	[]telegram.KeyboardButton{{Text: trashGoalCommand}, {Text: abortCommand}},
}

func (i *interaction) gotoWhatIsTheNextAction() {
	i.user.State = int(whatIsTheNextActionState)
	i.repo.update(i.user)
	i.sendPrompt("What is the next physical action?", whatIsTheNextActionKeyboard)
}

func (i *interaction) handleWhatIsTheNextAction() {
	i.gotoInitialState()
}
