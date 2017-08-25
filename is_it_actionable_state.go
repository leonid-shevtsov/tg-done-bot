package gtd_bot

import (
	"fmt"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

var isItActionableKeyboard = telegram.ReplyKeyboardMarkup{
	Keyboard: [][]telegram.KeyboardButton{
		[]telegram.KeyboardButton{{Text: yesCommand}, {Text: noCommand}, {Text: abortCommand}},
	},
	ResizeKeyboard:  true,
	OneTimeKeyboard: true,
}

func gotoProcessInbox(user *telegram.User) bool {
	return gotoIsItActionable(user)
}

func gotoIsItActionable(user *telegram.User) bool {
	var inboxItemsToProcess []InboxItem
	err := db.Model(&inboxItemsToProcess).
		Where("user_id = ? AND processed_at IS NULL", user.ID).
		Order("created_at ASC").
		Limit(1).
		Select()
	if err != nil {
		panic(err)
	}
	if len(inboxItemsToProcess) > 0 {
		inboxItemToProcess := inboxItemsToProcess[0]
		startProcessingInbox(user.ID, inboxItemToProcess.ID)
		responseText := fmt.Sprintf("*Processing item:*\n\n%s\n\n*Is this actionable?*", inboxItemToProcess.Text)
		msg := telegram.NewMessage(int64(user.ID), responseText)
		msg.ParseMode = telegram.ModeMarkdown
		msg.ReplyMarkup = isItActionableKeyboard
		bot.Send(msg)
		return true
	} else {
		bot.Send(telegram.NewMessage(int64(user.ID), "Inbox zero!"))
		return false
	}
}

func handleIsItActionableState(message *telegram.Message) {
	switch message.Text {
	case yesCommand:
		gotoWhatIsTheGoal(message.From)
	case noCommand:
		trashCurrentInboxItem(message.From)
	case abortCommand:
		abortProcessing(message.From)
	default:
		sendUnclearCommand(message.From, isItActionableKeyboard)
	}
}
