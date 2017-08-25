package gtd_bot

import (
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

func gotoWhatIsTheGoal(user *telegram.User) {
	setUserState(user.ID, whatIsTheGoalState)
	msg := telegram.NewMessage(int64(user.ID), "What is the goal here?")
	msg.ReplyMarkup = telegram.ReplyKeyboardMarkup{
		Keyboard: [][]telegram.KeyboardButton{
			[]telegram.KeyboardButton{{Text: trashGoalCommand}, {Text: abortCommand}},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}
	bot.Send(msg)
}

func handleWhatIsTheGoalState(message *telegram.Message) {
	switch message.Text {
	case trashGoalCommand:
		trashCurrentInboxItem(message.From)
	case abortCommand:
		abortProcessing(message.From)
	default:
		createGoal(message)
		gotoWhatIsTheNextAction(message.From)
	}
}

func createGoal(message *telegram.Message) {
	goal := Goal{
		UserID:    message.From.ID,
		Text:      message.Text,
		CreatedAt: time.Now(),
	}
	_, err := db.Model(&goal).Insert()
	if err != nil {
		panic(err)
	}
	_, err = db.Model(&UserState{
		UserID:        message.From.ID,
		CurrentGoalID: goal.ID,
	}).
		Set("current_goal_id = ?current_goal_id").
		Where("user_id = ?user_id").
		Update()
	if err != nil {
		panic(err)
	}
}
