package gtd_bot

import (
	"fmt"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

// FIXME
const processInboxCommand = "Process inbox"
const yesCommand = "Yes"
const noCommand = "No - trash it"
const abortCommand = "Let's do this later"
const trashGoalCommand = "Let's just trash it"

func abortProcessing(user *telegram.User) {
	abortText := fmt.Sprintf("OK. Items left to process: %d. Feel free to brain dump.", inboxCount(user.ID))
	gotoBackToInitialState(user, abortText)
}

func sendUnclearCommand(user *telegram.User, keyboardMarkup telegram.ReplyKeyboardMarkup) {
	msg := telegram.NewMessage(int64(user.ID), "Pick one of the options, please.")
	msg.ReplyMarkup = keyboardMarkup
	bot.Send(msg)
}

func trashCurrentInboxItem(user *telegram.User) {
	inboxItem := getCurrentInboxItem(user.ID)
	inboxItem.ProcessedAt = time.Now()
	_, err := db.Model(inboxItem).Update()
	if err != nil {
		panic(err)
	}
	bot.Send(telegram.NewMessage(int64(user.ID), "Trashed! Moving on."))
	if !gotoProcessInbox(user) {
		gotoBackToInitialState(user, "Feel free to brain dump.")
	}
}

func inboxCount(userID int) int {
	count, err := db.Model(&InboxItem{}).Where("user_id = ? AND processed_at IS NULL", userID).Count()
	if err != nil {
		panic(err)
	}
	return count
}

func updateLastMessageAt(userID int) {
	userState := UserState{UserID: userID, LastMessageAt: time.Now()}
	_, err := db.Model(&userState).
		Where("user_id = ?user_id").
		OnConflict("(user_id) DO UPDATE").
		Set("last_message_at = ?last_message_at").
		Insert()
	if err != nil {
		panic(err)
	}
}

func startProcessingInbox(userID int, inboxItemID int) {
	_, err := db.Model(&UserState{
		UserID:             userID,
		CurrentInboxItemID: inboxItemID,
		State:              int(isItActionableState),
	}).
		Set("current_inbox_item_id = ?current_inbox_item_id, state = ?state").
		Where("user_id = ?user_id").
		Update()
	if err != nil {
		panic(err)
	}
}

func getUserState(userID int) interactionState {
	var userState UserState
	err := db.Model(&userState).Where("user_id = ?", userID).First()
	if err != nil {
		panic(err)
	}
	return interactionState(userState.State)
}

func getCurrentInboxItem(userID int) *InboxItem {
	var userState UserState
	err := db.Model(&userState).Column("CurrentInboxItem").First()
	if err != nil {
		panic(err)
	}
	return userState.CurrentInboxItem
}

func setUserState(userID int, newState interactionState) {
	_, err := db.Model(&UserState{UserID: userID, State: int(newState)}).
		Where("user_id = ?user_id").
		Set("state = ?state").
		Update()
	if err != nil {
		panic(err)
	}
}
