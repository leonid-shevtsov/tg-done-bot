package gtd_bot

import telegram "github.com/go-telegram-bot-api/telegram-bot-api"

func gotoWhatIsTheNextAction(user *telegram.User) {
	setUserState(user.ID, whatIsTheNextActionState)
	msg := telegram.NewMessage(int64(user.ID), "What is the next physical action?")
	msg.ReplyMarkup = telegram.ReplyKeyboardMarkup{
		Keyboard: [][]telegram.KeyboardButton{
			[]telegram.KeyboardButton{{Text: trashGoalCommand}, {Text: abortCommand}},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}
	bot.Send(msg)
}
