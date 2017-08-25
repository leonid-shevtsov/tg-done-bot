package gtd_bot

import (
	"fmt"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

var initialStateKeyboard = telegram.ReplyKeyboardMarkup{
	Keyboard: [][]telegram.KeyboardButton{
		[]telegram.KeyboardButton{{Text: processInboxCommand}},
	},
	ResizeKeyboard:  true,
	OneTimeKeyboard: true,
}

func gotoBackToInitialState(user *telegram.User, message string) {
	setUserState(user.ID, initialState)
	msg := telegram.NewMessage(int64(user.ID), message)
	if inboxCount(user.ID) > 0 {
		msg.ReplyMarkup = initialStateKeyboard
	}
	bot.Send(msg)
}

func handleInitialState(message *telegram.Message) {
	switch message.Text {
	case processInboxCommand:
		gotoProcessInbox(message.From)
	default:
		addInboxItem(message.From, message.Text)
	}
}

func addInboxItem(user *telegram.User, text string) {
	inboxItem := InboxItem{
		UserID:    user.ID,
		Text:      text,
		CreatedAt: time.Now(),
	}
	db.Insert(&inboxItem)

	responseText := fmt.Sprintf("Added to inbox. Now in inbox: %d items.", inboxCount(user.ID))
	msg := telegram.NewMessage(int64(user.ID), responseText)
	msg.ReplyMarkup = initialStateKeyboard
	bot.Send(msg)
}
