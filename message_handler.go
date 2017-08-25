package gtd_bot

import telegram "github.com/go-telegram-bot-api/telegram-bot-api"

func handleMessage(message *telegram.Message) {
	user := message.From
	updateLastMessageAt(user.ID)
	state := getUserState(user.ID)

	switch state {
	case initialState:
		handleInitialState(message)
	case isItActionableState:
		handleIsItActionableState(message)
	case whatIsTheGoalState:
		handleWhatIsTheGoalState(message)
	}
}
